package main

import (
	"fmt"
	"math"
)

func main() {

	var num int
	fmt.Println("Enter four-digit number")
	fmt.Scan(&num)

	digits := []int{}

	for num > 0 {
		nums := num % 10
		digits = append(digits, (nums+7)%10)
		num /= 10
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	digits[0], digits[2] = digits[2], digits[0]
	digits[1], digits[3] = digits[3], digits[1]

	result := 0
	for i, digit := range digits {
		result += digit * int(math.Pow10(len(digits)-i-1))
	}

	fmt.Println(result)
}
