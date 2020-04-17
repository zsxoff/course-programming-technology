package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"strconv"
	"strings"
)

func inRange(val int, arr *[]int) bool {
	for _, x := range *arr {
		if x == val {
			return true
		}
	}
	return false
}

func readInput(variations *[]int) int {
	// Generate help message.
	var vals []string

	for _, x := range *variations {
		vals = append(vals, strconv.Itoa(x))
	}

	if len(vals) == 0 {
		log.Fatal("No values in read input.")
	}

	help := "введите " + strings.Join(vals[:len(vals)-1], ", ") + " или " + vals[len(vals)-1] + ": "

	// Scan value.
	var (
		val int
		err error
	)

	for {
		fmt.Print(help)

		// Scan value.
		var scanLine string
		_, err = fmt.Scanln(&scanLine)
		if err != nil {
			break
		}

		// Check if value is integer and in range.
		val, err = strconv.Atoi(scanLine)
		if err == nil {
			if inRange(val, variations) {
				break
			}
		}

		color.Red("\n! Ошибка: Неизвестное значение <" + scanLine + "> !\n")
	}

	fmt.Println()

	return val
}
