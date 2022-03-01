package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := lire("Enter the N Number to count to : ")

	var result string
	for i := 0; i < num; i++ {
		result += fmt.Sprintf("%d ", i)
	}

	fmt.Println(strings.TrimSpace(result))
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
