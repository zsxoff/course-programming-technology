package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"log"
	"strconv"
)

// Player actions status.
const (
	ActionChill  = iota
	ActionAttack = iota
)

// Player users structure.
type Player struct {
	Nickname string `json:"nickname"` // Username
	Warriors int    `json:"warriors"` // Warriors count.
	Action   int    `json:"action"`   // Action - chill or attack.
	Crystals int    // Crystals count.
	Workers  int    // Workers count.
	turn     int    // Current turn number.
}

// addCrystals function adds the number of crystals equal to the number of workers.
func (p *Player) addCrystals() {
	p.Crystals += p.Workers
}

// NextTurn function apply addCrystals() and increase turns count by 1.
func (p *Player) NextTurn() {
	p.turn++
	p.addCrystals()
}

// HireWorker function adds the one worker and subtracts 5 crystals.
func (p *Player) HireWorker() bool {
	if (p.Crystals) < 5 {
		return false
	}
	p.Workers++
	p.Crystals -= 5
	return true
}

// HireWarrior function adds the one warrior and subtracts 10 crystals.
func (p *Player) HireWarrior() bool {
	if (p.Crystals) < 10 {
		return false
	}
	p.Warriors++
	p.Crystals -= 10
	return true
}

// PrintCurrentTurn print colored current turn number.
func (p *Player) PrintCurrentTurn() {
	color.Green("\n-= Turn #" + strconv.Itoa(p.turn) + " =-\n\n")
}

// PrintResources function print player's current resources.
func (p *Player) PrintResources() {
	color.Yellow("You have " +
		strconv.Itoa(p.Crystals) + " crystals, " +
		strconv.Itoa(p.Workers) + " workers and " +
		strconv.Itoa(p.Warriors) + " warriors.\n\n")
}

// MakeDecision function create dialog with player and run functions.
func (p *Player) MakeDecision() {

	// Read action.
	p.PrintCurrentTurn()
	p.PrintResources()

	fmt.Println("Actions:" +
		"\n" +
		"\n[1] Hire workers or warriors" +
		"\n[2] Attack another player")
	fmt.Println()

	action := readInput(&[]int{1, 2})

	switch action {

	// Hire menu.
	case 1:
		p.Action = ActionChill
		for true {
			p.PrintResources()

			color.Cyan("Whom to hire?")
			fmt.Println("\n[1] Worker (cost 5 crystals)" +
				"\n[2] Warrior (cost 10 crystals)" +
				"\n[3] Exit and end of turn")
			fmt.Println()

			hire := readInput(&[]int{1, 2, 3})

			// Exit.
			if hire == 3 {
				break
			}

			// Hire type.
			switch hire {
			case 1:
				if p.HireWorker() {
					color.Green("You hired 1 worker.")
				} else {
					color.Red("Error: You don't have enough crystals.")
				}
			case 2:
				if p.HireWarrior() {
					color.Green("You hired 1 warrior.")
				} else {
					color.Red("Error: You don't have enough crystals.")
				}
			}
		}

	// Attack menu.
	case 2:
		p.Action = ActionAttack
	}

	color.Green("-= End of turn =-")
	p.NextTurn()
}

// Init function clean up player except nickname.
func (p *Player) Init() {
	p.Warriors = 0
	p.Action = ActionChill
	p.Crystals = 5
	p.Workers = 0
	p.turn = 0
}

// ToJson function convert player to string JSON.
func (p *Player) ToJson() string {
	bytesJson, err := json.Marshal(p)

	if err != nil {
		log.Fatal(err)
	}

	return string(bytesJson)
}

// FromJson convert string JSON to player.
func (p *Player) FromJson(stringJson string) {
	err := json.Unmarshal([]byte(stringJson), &p)

	if err != nil {
		log.Println(err)
	}
}
