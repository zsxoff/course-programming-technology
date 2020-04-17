package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
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

// PrintTurnBegin print colored begin text and current turn number.
func (p *Player) PrintTurnBegin() {
	color.Green("\n= = = = = НАЧАЛО ХОДА " + strconv.Itoa(p.Turn) + " = = = = =\n\n")
}

// PrintTurnEnd print colored end text and current turn number.
func (p *Player) PrintTurnEnd() {
	color.Green("\n= = = = = КОНЕЦ ХОДА " + strconv.Itoa(p.Turn) + " = = = = =\n\n")
}

// PrintResources function print player's current resources.
func (p *Player) PrintResources() {
	color.Yellow("Ваши ресурсы: " +
		strconv.Itoa(p.CountCrystals) + " кристаллов, " +
		strconv.Itoa(p.CountWorkers) + " рабочих и " +
		strconv.Itoa(p.CountWarriors) + " воинов.\n\n")
}

// MakeDecision function create dialog with player and run functions.
func (p *Player) MakeDecision() {

	// Read action.
	p.PrintTurnBegin()
	p.PrintResources()

	fmt.Println("Действия:" +
		"\n" +
		"\n[1] Нанять рабочих или воинов" +
		"\n[2] Атаковать другого игрока")
	fmt.Println()

	action := readInput(&[]int{1, 2})

	switch action {

	// Hire menu.
	case 1:
		p.Action = ActionChill
		for true {
			p.PrintResources()

			color.Cyan("Кого нанять?")
			fmt.Println("\n[1] Рабочий (цена 5 кристаллов)" +
				"\n[2] Воин (цена 10 кристаллов)" +
				"\n[3] Выйти и закончить ход")
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
					color.Green("Вы наняли 1 рабочего.")
				} else {
					color.Red("У вас не хватает кристаллов для найма рабочего!")
				}
			case 2:
				if p.HireWarrior() {
					color.Green("ВЫ наняли 1 воина.")
				} else {
					color.Red("У вас не хватает кристаллов для найма воина!")
				}
			}
		}

	// Attack menu.
	case 2:
		p.Action = ActionAttack
	}

	p.PrintTurnEnd()

	// Increase player resources.
	p.NextTurn()

	fmt.Println()
	color.Blue("Ожидание другого игрока...")
	fmt.Println()
}

// Init function clean up player except nickname.
func (p *Player) Init() {
	p.CountWarriors = 0
	p.Action = ActionChill
	p.CountCrystals = 5
	p.CountWorkers = 0
	p.Turn = 1
}

// ToJson function convert player to string JSON.
func (p *Player) ToJson() (string, error) {
	bytesJson, err := json.Marshal(p)

	if err != nil {
		return "", err
	}

	return string(bytesJson), nil
}

// FromJson convert string JSON to player.
func (p *Player) FromJson(stringJson *string) error {
	return json.Unmarshal([]byte(*stringJson), &p)
}
