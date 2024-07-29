package main

import "fmt"

func main() {
	var largest int
	var number int
	fmt.Scan(largest)

	for count := 1; count <= 10; count++ {
		fmt.Println("Enter contestant value")
		fmt.Scan(&number)

		if number > largest {
			largest = number

		}

	}
	fmt.Println("Largest number is", largest)

}
