package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := lire("Divisible par 7 et 13 : ")
	imprimeDiv7ou13(num)
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

func imprimeDiv7ou13(n int) {
	for curr := 1; curr <= n; curr++ {
		if curr%7 == 0 && curr%13 == 0 {
			fmt.Println(curr)
		}
	}
}
