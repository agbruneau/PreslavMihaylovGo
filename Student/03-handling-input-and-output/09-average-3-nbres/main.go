package main

import "fmt"

func main() {
	/*	var num1 int
		fmt.Scanln(&num1)
		var num2 float64
		fmt.Scanln(&num2)
		fmt.Println(num1, num2)
	*/
	var num1, num2, num3, sum float64
	fmt.Print("Saisir le premier nombre (1/3) : ")
	fmt.Scanln(&num1)

	fmt.Print("Saisir le second nombre de (2/3) : ")
	fmt.Scanln(&num2)

	fmt.Print("Saisir le troisième nombre de (3/3) : ")
	fmt.Scanln(&num3)

	sum = num1 + num2 + num3
	fmt.Printf("%f + %f + %f = %f\n", num1, num2, num3, sum)
}
