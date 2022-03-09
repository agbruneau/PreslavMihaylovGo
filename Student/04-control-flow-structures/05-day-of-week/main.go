package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	dayOfWeekIndex := lire("Day number of the week : ")
	dayOfWeek(dayOfWeekIndex)
}

func lire(s string) int {
	var line string
	fmt.Print(s)

	if _, err := fmt.Scanln(&line); err != nil {
		panic(err)
	}

	n, err := strconv.Atoi(strings.TrimSpace(line))

	if err != nil {
		panic(err)
	}

	return n
}

func dayOfWeek(index int) {
	switch index {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Day of the week not recognized")
	}

}
