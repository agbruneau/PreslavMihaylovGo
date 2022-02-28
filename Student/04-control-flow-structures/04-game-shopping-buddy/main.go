package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := lire("Money on hand to buy games ? ")
	gamesBought := []string{}

	if num >= 30 {
		gamesBought = append(gamesBought, "StarCraft 2")
		num -= 30
	}

	if num >= 130 {
		gamesBought = append(gamesBought, "Cyberpunk 2077")
		num -= 130
	}

	if num >= 60 {
		gamesBought = append(gamesBought, "Witcher 3")
		num -= 60
	}

	if len(gamesBought) == 0 {
		fmt.Println("I couldn't buy anything!")
	} else {
		fmt.Println("Here's what I bought : ")
		fmt.Println(strings.Join(gamesBought, ", "))
	}
}

func lire(s string) int {
	var line string
	fmt.Print(s)
	fmt.Scanln(&line)
	n, err := strconv.Atoi(strings.TrimSpace(line))
	erreur(err)
	return n
}

func erreur(err error) {
	if err != nil {
		panic(err)
	}
}
