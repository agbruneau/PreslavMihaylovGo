package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	rangeOfNumbers := lire("Range of number to guess : ")

	userGuess := lire("Your guess is : ")

	myGuess := tirageAuSort(rangeOfNumbers)

	verifier(userGuess, myGuess)

}

func lire(s string) int {
	var l string
	fmt.Print(s)
	fmt.Scanln(&l)
	n, err := strconv.Atoi(strings.TrimSpace(l))
	erreur(err)
	return n
}

func tirageAuSort(limite int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(limite)
}

func verifier(n1 int, n2 int) {
	if n1 == n2 {
		fmt.Println("You guessed the number correctrly!")
	} else {
		fmt.Println("That's not the number I had in mind.")
		fmt.Printf("You guessed %d but my number was %d\n", n1, n2)
	}

}

func erreur(err error) {
	if err != nil {
		panic(err)
	}
}
