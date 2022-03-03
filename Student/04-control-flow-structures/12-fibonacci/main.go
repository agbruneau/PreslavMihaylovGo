package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := lire("Calcul Fibonacci : ")
	imprimeFibonacci(num)
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

func imprimeFibonacci(n int) {
	premier, second := 1, 1
	fmt.Printf("%d %d ", premier, second)

	for i := 3; i <= n; i++ {
		curr := premier + second
		premier = second
		second = curr
		fmt.Printf("%d ", curr)
	}
}
