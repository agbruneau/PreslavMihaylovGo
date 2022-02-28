package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var line string

	fmt.Printf("Read Number one :")
	fmt.Scanln(&line)
	num1, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read Number two :")
	fmt.Scanln(&line)
	num2, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	if num1 > num2 {
		fmt.Printf("The bigger number is %d\n", num1)
	} else if num1 < num2 {
		fmt.Printf("The bigger number is %d\n", num2)
	} else {
		fmt.Printf("The number are equal\n")
	}

}
