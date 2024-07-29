package main

import "fmt"

func main() {
	var largest int
	var number int
	var secondLargest int

	for count := 1; count < 10; count++ {
		fmt.Println("Enter number: ")
		fmt.Scanf("%d", &number)

	}
	if number > largest {
		largest = number
	}
	if number != largest {
		secondLargest = number
	}

	fmt.Printf("Largest number: %d\n", largest)
	fmt.Printf("Second largest number: %d\n", secondLargest)
}
