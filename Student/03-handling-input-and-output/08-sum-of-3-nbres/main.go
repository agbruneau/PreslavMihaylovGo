package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	validation(err)

	num1, err := strconv.Atoi(strings.TrimSpace(line))
	validation(err)

	line, err = r.ReadString('\n')
	validation(err)

	num2, err := strconv.Atoi(strings.TrimSpace(line))
	validation(err)

	line, err = r.ReadString('\n')
	validation(err)

	num3, err := strconv.Atoi(strings.TrimSpace(line))
	validation(err)

	fmt.Println(num1 + num2 + num3)
}

func validation(err error) {
	if err != nil {
		panic(err)
	}
}
