package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := lire("Calcul du factoriel : ")
	fac := factoriel(num)
	fmt.Printf("Résultat : %d", fac)
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

func factoriel(n int) int {
	fac := 1
	for i := 1; i <= n; i++ {
		fac *= i
	}
	return fac
}
