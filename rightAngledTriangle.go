package main

import "fmt"

func main() {
	var asterisk int
	fmt.Print("Enter your asterisk: ")
	fmt.Scan(&asterisk)

	if asterisk < 1 || asterisk > 10 {

		fmt.Print("Out of Range")
		return
	}

	for count := 1; count <= asterisk; count++ {
		for counter := 1; counter <= count; counter++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}
