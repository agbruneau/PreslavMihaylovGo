package main

import (
	"fmt"
	"strings"
)

func main() {
	character := lire("Name of the character : ")
	var game string
	if isCharacterFromTheWitcher3(character) {
		game = "The Witcher 3"
	} else if isCharacterFromStarcraft2(character) {
		game = "Starcraft 2"
	} else if isCharacterFromWarCraft3(character) {
		game = "WarCraft 3"
	}

	fmt.Printf("%s is part of the game %s!\n", character, game)

}

func lire(s string) string {
	var l string
	fmt.Print(s)
	fmt.Scanln(&l)
	c := strings.TrimSpace(l)
	return c
}

func isCharacterFromTheWitcher3(character string) bool {
	return character == "Geralt" ||
		character == "Ciri" ||
		character == "Yennefer"
}

func isCharacterFromStarcraft2(character string) bool {
	return character == "Aldaris" ||
		character == "Jim Raynor" ||
		character == "Kerrigan" ||
		character == "Zeratul"
}

func isCharacterFromWarCraft3(character string) bool {
	return character == "Arthas" ||
		character == "Illidan" ||
		character == "Sylvanas"
}
