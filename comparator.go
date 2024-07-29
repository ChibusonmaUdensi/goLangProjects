package main

import "fmt"

func main() {
	fmt.Println("Enter first number")
	var firstNumber int
	fmt.Scanf("%d", &firstNumber)

	fmt.Println("Enter second number")
	var secondNumber int
	fmt.Scanf("%d", &secondNumber)

	if firstNumber == secondNumber {
		fmt.Println("0")
	}
	if firstNumber > secondNumber {
		fmt.Println("1")
	}
	if firstNumber < secondNumber {
		fmt.Println("-1")
	}
}
