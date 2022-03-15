package main

import "fmt"

/*
Go's basic types are

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte alias for uint8

rune alias for int32

float32 float64

complex64 complex128

*/

func main() {
	var n1 float32 = 14.435234
	var n2 float64 = 14.435234234234235
	fmt.Println(n1, n2)
}
