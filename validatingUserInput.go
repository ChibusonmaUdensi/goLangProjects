package main

import "fmt"

func main() {
	var input int

	for input > 2 {
		fmt.Println("Please enter a valid number")
		fmt.Scan(&input)
	}

	fmt.Println("Valid Input")
}
