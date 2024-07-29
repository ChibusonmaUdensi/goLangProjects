package main

import (
	"fmt"
	"math"
)

func binaryToDecimal(num int) int {
	var remainder int
	index := 0
	decimalNum := 0

	for num != 0 {
		remainder = num % 10
		num = num / 10
		decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
		index++
	}
	return decimalNum
}
func main() {
	fmt.Println("Enter binary number:")
	var binary int
	fmt.Scan(&binary)
	fmt.Println(binaryToDecimal(binary))
}
