package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	StatusW         = "W"
	StatusF         = "F"
	StatusD         = "D"
	StatusUndefined = "U"
)

type ConnectionConfig struct {
	Ip   string
	Port int
}

func messageSend(text *string, conn *net.Conn) {
	_, err := fmt.Fprintf(*conn, *text+"\n")
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

func startServerMode(connConfig *ConnectionConfig) {
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Try to create server.
	fmt.Print("Создание нового сервер... ")

	lisn, err := net.Listen("tcp4", ":"+strconv.Itoa(connConfig.Port))
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
		defer conn.Close()

		// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
		// Cleanup player.
		fmt.Print("Создание нового персонажа... ")

		// Init player.
		serverPlayer := Player{}
		serverPlayer.Init()

		fmt.Println("ОК")
		// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

		// Run game.
		fmt.Println("Игра началась!")
		for {
			serverPlayer.MakeDecision()

			// Receive message.
			messageRecv := messageRecv(&conn)

			// Create client player entity.
			var clientPlayer Player

			err := clientPlayer.FromJson(&messageRecv)
			if err != nil {
				log.Fatal(err)
			}

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
			messageSend(&playerStatus, &conn)

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

		fmt.Println()
		fmt.Println("Ожидание нового игрока (вы можете остановить сервер с помощью Ctrl+C)")
	}
}

func startClientMode(connConfig *ConnectionConfig) {
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Try to connect.
	fmt.Print("Попытка подключиться к удаленному серверу... ")

	address := connConfig.Ip + ":" + strconv.Itoa(connConfig.Port)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("ОК")
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Cleanup player.
	fmt.Print("Создание нового персонажа... ")

	p := Player{}
	p.Init()

	fmt.Println("ОК")
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Run game.
	fmt.Println("Игра началась!")
	for {
		p.MakeDecision()

		// Send player status.
		marshal, err := p.ToJson()
		if err != nil {
			log.Fatal(err)
		}
		messageSend(&marshal, &conn)

		// Receive message.
		playerStatus := messageRecv(&conn)

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

	fmt.Println()
	fmt.Println("Вы можете начать новую игру снова подключившись к серверу.")
}

func main() {
	// Parse args.
	var (
		mode string
		ip   string
		port int
	)

	flag.StringVar(&mode, "mode", "undefined", "client or server mode")
	flag.StringVar(&ip, "ip", "undefined", "IP address of server")
	flag.IntVar(&port, "port", -1, "port of server")
	flag.Parse()

	// Check args.
	switch mode {
	case "client":
		if ip == "undefined" {
			log.Fatal("Error: incorrect IP address!")
		}
		if port == -1 {
			log.Fatal("Error: incorrect port number!")
		}
	case "server":
		if port == -1 {
			log.Fatal("Error: incorrect port number!")
		}
	}

	// Init connection structure.
	connConfig := ConnectionConfig{Ip: ip, Port: port}

	// Init connection.
	switch mode {
	case "client":
		startClientMode(&connConfig)
	case "server":
		startServerMode(&connConfig)
	default:
		color.Red("Неизвестный режим игры! Используйте параметр client или server")
	}
}
