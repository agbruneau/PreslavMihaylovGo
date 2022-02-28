package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	x := lire("Number to validate in range : ")
	a := lire("Number lower range : ")
	b := lire("Number upper range : ")
	if x >= a && x <= b {
		fmt.Printf("%d is in range [%d, %d]", x, a, b)
	} else {
		fmt.Printf("%d is not in range [%d, %d]", x, a, b)
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
