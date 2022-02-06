package main

import "fmt"

func main() {
	// a := 14
	// b := 32
	a, b := 14, 32
	imprime(a, b)
	b, a = a, b
	imprime(a, b)
}

func imprime(a int, b int) {
	fmt.Printf("a = %d\nb = %d\n", a, b)
}
