package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := lire("Vérification des nombres premiers : ")
	imprimeNbrPremier(num)
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

func imprimeNbrPremier(n int) {
	for curr := 0; curr <= n; curr++ {
		isPrime := true
		for dividedBy := 2; dividedBy < curr; dividedBy++ {
			if curr%dividedBy == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Printf("%d ", curr)
		}

	}
}
