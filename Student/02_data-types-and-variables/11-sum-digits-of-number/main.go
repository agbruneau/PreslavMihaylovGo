package main

import "fmt"

func main() {
	sum := 0
	num := 123456789

	for num > 0 {
		digit := num % 10
		num /= 10
		sum += digit
	}
	fmt.Println(sum)
}
