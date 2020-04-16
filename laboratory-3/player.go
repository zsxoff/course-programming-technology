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
	Action        int `json:"action"`        // Action - chill or attack.
	Turn          int `json:"turn"`          // Current turn number.
	CountWorkers  int `json:"countWorkers"`  // Workers count.
	CountCrystals int `json:"countCrystals"` // Crystals count.
	CountWarriors int `json:"countWarriors"` // Warriors count.
}

// addCrystals function adds the number of crystals equal to the number of workers.
func (p *Player) addCrystals() {
	p.CountCrystals += p.CountWorkers
}

// NextTurn function apply addCrystals() and increase turns count by 1.
func (p *Player) NextTurn() {
	p.Turn++
	p.addCrystals()
}

// HireWorker function adds the one worker and subtracts 5 crystals.
func (p *Player) HireWorker() bool {
	if (p.CountCrystals) < 5 {
		return false
	}
	p.CountWorkers++
	p.CountCrystals -= 5
	return true
}

// HireWarrior function adds the one warrior and subtracts 10 crystals.
func (p *Player) HireWarrior() bool {
	if (p.CountCrystals) < 10 {
		return false
	}
	p.CountWarriors++
	p.CountCrystals -= 10
	return true
}

// PrintCurrentTurn print colored current Turn number.
func (p *Player) PrintCurrentTurn() {
	color.Green("\n-= Turn #" + strconv.Itoa(p.Turn) + " =-\n\n")
}

// PrintResources function print player's current resources.
func (p *Player) PrintResources() {
	color.Yellow("You have " +
		strconv.Itoa(p.CountCrystals) + " crystals, " +
		strconv.Itoa(p.CountWorkers) + " workers and " +
		strconv.Itoa(p.CountWarriors) + " warriors.\n\n")
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
				"\n[3] Exit and end of Turn")
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

	color.Green("-= End of Turn =-")
	p.NextTurn()
}

// Init function clean up player except nickname.
func (p *Player) Init() {
	p.CountWarriors = 0
	p.Action = ActionChill
	p.CountCrystals = 5
	p.CountWorkers = 0
	p.Turn = 0
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
