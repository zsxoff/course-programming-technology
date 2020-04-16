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

func startServerMode(p *Player) {
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Cleanup player.
	fmt.Print("Создание нового персонажа... ")

	p.Init()

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
			p.MakeDecision()

			// Receive message.
			messageRecv := messageRecv(&conn)

			// Create client player entity.
			var clientPlayer Player
			clientPlayer.FromJson(messageRecv)

			// Resolve players battle.
			playerStatus := StatusUndefined

			if clientPlayer.Action == ActionAttack || p.Action == ActionAttack {
				if clientPlayer.Warriors > p.Warriors {
					playerStatus = StatusW // Player win.
				} else {
					playerStatus = StatusF // Player lose.
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
		messageRecv := messageRecv(&conn)
		log.Println("recv " + messageRecv)

		// Process message.
		if messageRecv == StatusW {
			color.Green("Ура, вы победили!")
			break
		} else if messageRecv == StatusF {
			color.Red("Увы, вы проиграли! У вас не хватило воинов для атаки/защиты...")
			break
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
	fmt.Print("Enter your nickname: ")

	var nickname string
	_, err := fmt.Scanln(&nickname)
	if err != nil {
		log.Fatal(err)
	}
	color.Cyan("\nHello, " + nickname + "!")

	player := Player{Crystals: 25, Workers: 5, Warriors: 0}
	player.Nickname = nickname

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
