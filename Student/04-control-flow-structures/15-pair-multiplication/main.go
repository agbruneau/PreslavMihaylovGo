package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := lire("Saisir N : ")
	x := lire("Saisir X : ")
	imprimeMulParPair(n, x)
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

func imprimeMulParPair(n int, x int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i*j == x {
				fmt.Printf("%d * %d = %d\n", i, j, x)
			}
		}
	}
}
