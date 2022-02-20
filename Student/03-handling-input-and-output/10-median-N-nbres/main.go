package main

import "fmt"

func main() {
	var cnt int
	nums := []int{}
	num, median := 0, 0

	fmt.Print("Quantité de nombres à saisir : ")
	fmt.Scanln(&cnt)

	for i := 0; i < cnt; i++ {
		fmt.Printf("Saisir le %d de %d : ", i, cnt)
		fmt.Scanln(&num)
		nums = append(nums, num)
	}

	median = nums[cnt/2]
	fmt.Printf("Median = %d\n", median)

}
