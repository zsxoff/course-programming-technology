package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"log"
	"net"
	"strings"
)

const PORT = ":7778"

const (
	StatusW         = "W"
	StatusF         = "F"
	StatusD         = "D"
	StatusUndefined = "U"
)

func messageSend(text string, conn *net.Conn) {
	_, err := fmt.Fprintf(*conn, text+"\n")
	if err != nil {
		log.Fatal(err)
	}
}

func messageRecv(conn *net.Conn) string {
	messageRecv, err := bufio.NewReader(*conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(messageRecv)
}

func startServerMode(serverPlayer *Player) {
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Cleanup player.
	fmt.Print("Создание нового персонажа... ")

	serverPlayer.Init()

	fmt.Println("ОК")
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Try to create server.
	fmt.Print("Создание нового сервер... ")

	lisn, err := net.Listen("tcp4", PORT)
	defer lisn.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ОК")
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	color.Green("Теперь вы - сервер! Вы можете пригласить соперника по адресу: " + lisn.Addr().String())

	// Accept new connection.
	fmt.Println("Ожидание игроков...")
	for {
		conn, err := lisn.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Run game.
		fmt.Println("Игра началась!")
		for {
			serverPlayer.MakeDecision()

			// Receive message.
			messageRecv := messageRecv(&conn)

			// Create client player entity.
			var clientPlayer Player
			clientPlayer.FromJson(messageRecv)

			// Resolve players battle.
			playerStatus := StatusUndefined

			// If attacker - Client.
			if clientPlayer.Action == ActionAttack && serverPlayer.Action == ActionChill {
				if clientPlayer.CountWarriors > serverPlayer.CountWarriors {
					playerStatus = StatusW // Player win.
				} else {
					playerStatus = StatusF // Player lose.
				}
			}

			// If attacker - Server.
			if serverPlayer.Action == ActionAttack && clientPlayer.Action == ActionChill {
				if serverPlayer.CountWarriors > clientPlayer.CountWarriors {
					playerStatus = StatusF // Server win.
				} else {
					playerStatus = StatusW // Server lose.
				}
			}

			// If attacker - Client and Server.
			if clientPlayer.Action == ActionAttack && serverPlayer.Action == ActionAttack {
				if clientPlayer.CountWarriors > serverPlayer.CountWarriors {
					playerStatus = StatusW // Player win.
				} else if clientPlayer.CountWarriors < serverPlayer.CountWarriors {
					playerStatus = StatusF // Player lose.
				} else {
					playerStatus = StatusD // Draw.
				}
			}

			// Send message.
			messageSend(playerStatus, &conn)

			// Process message.
			if playerStatus == StatusF {
				color.Green("Ура, вы победили!")
				break
			} else if playerStatus == StatusW {
				color.Red("Увы, вы проиграли! У вас не хватило воинов для атаки/защиты...")
				break
			} else if playerStatus == StatusD {
				color.Yellow("Игроки напали одновременно, но не смогли победить друг друга!")
			} else {
				color.Yellow("За прошедший ход не было сражений...")
			}
		}

		fmt.Println("Ожидание нового игрока (вы можете остановить сервер с помощью Ctrl+C)")
	}
}

func startClientMode(p *Player) {

	CONNECT := "127.0.0.1" + PORT

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Cleanup player.
	fmt.Print("Создание нового персонажа... ")

	p.Init()

	fmt.Println("ОК")
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Try to connect.
	fmt.Print("Попытка подключиться к удаленному серверу... ")

	conn, err := net.Dial("tcp", CONNECT)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ОК")
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Run game.
	fmt.Println("Игра началась!")
	for {
		p.MakeDecision()

		// Send player status.
		messageSend(p.ToJson(), &conn)

		// Receive message.
		playerStatus := messageRecv(&conn)
		log.Println("recv " + playerStatus)

		// Process message.
		if playerStatus == StatusW {
			color.Green("Ура, вы победили!")
			break
		} else if playerStatus == StatusF {
			color.Red("Увы, вы проиграли! У вас не хватило воинов для атаки/защиты...")
			break
		} else if playerStatus == StatusD {
			color.Yellow("Игроки напали одновременно, но не смогли победить друг друга!")
		} else {
			color.Yellow("За прошедший ход не было сражений...")
		}
	}
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	// Close connection.
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var (
		mode string
	)

	flag.StringVar(&mode, "mode", "undefined", "client or server mode")
	flag.Parse()

	// Init player.
	player := Player{CountCrystals: 25, CountWorkers: 5, CountWarriors: 0}

	// Init connection.
	switch mode {
	case "client":
		startClientMode(&player)
	case "server":
		startServerMode(&player)
	default:
		color.Red("Неизвестный режим игры! Используйте параметр client или server")
	}
}
