package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nombre := lire("Enter number to translate in word : ")
	nombretxt := convertirNombreEnTexte(nombre)
	fmt.Println(nombretxt)
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

func convertirNombreEnTexte(num int) string {
	hundreds := num / 100

	var result string
	switch hundreds {
	case 1:
		result += "one hundred"
	case 2:
		result += "two hundred"
	case 3:
		result += "three hundred"
	case 4:
		result += "four hundred"
	case 5:
		result += "five hundred"
	case 6:
		result += "six hundred"
	case 7:
		result += "seven hundred"
	case 8:
		result += "eight hundred"
	case 9:
		result += "nine hundred"
	}

	if hundreds > 0 && num%100 > 0 {
		result += " and "
	}

	result += formatTens(num % 100)
	return strings.TrimSpace(result)
}

func formatTens(num int) string {
	tens := num / 10
	var result string
	switch tens {
	case 1:
		return formatSpecialTens(num)
	case 2:
		result += "twenty "
	case 3:
		result += "thirty "
	case 4:
		result += "fourty "
	case 5:
		result += "fifty "
	case 6:
		result += "sixty "
	case 7:
		result += "seventy "
	case 8:
		result += "eighty "
	case 9:
		result += "ninety "
	}

	result += formatOnes(num % 10)
	return result
}

func formatSpecialTens(num int) string {
	switch num {
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "fifteen"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	default:
		return ""
	}
}

func formatOnes(num int) string {
	var result string
	switch num {
	case 1:
		result += "one"
	case 2:
		result += "two"
	case 3:
		result += "three"
	case 4:
		result += "four"
	case 5:
		result += "five"
	case 6:
		result += "six"
	case 7:
		result += "seven"
	case 8:
		result += "eight"
	case 9:
		result += "nine"
	}

	return result
}
