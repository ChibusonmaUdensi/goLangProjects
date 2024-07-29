package main

import "fmt"

func main() {
	for count := 1; count <= 8; count++ {
		for counter := 1; counter <= 16; counter++ {
			if (count+counter)%2 == 0 {
				fmt.Print("* ")

			}

			fmt.Print(" ")
		}
		fmt.Println()
	}

}
