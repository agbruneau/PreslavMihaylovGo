package main

import "fmt"

func main() {
	// Input value
	var input1 float32 = 3.14
	var input2 float64 = 12.3456789
	var input3 int16 = 1234
	var input4 int16 = 1234

	// Output value
	var output1 int = int(input1)
	var output2 float32 = float32(input2)
	var output3 int8 = int8(input3)
	var output4 uint8 = uint8(input4)

	fmt.Println(output1, output2, output3, output4)
}
