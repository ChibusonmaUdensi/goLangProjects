package main

import "fmt"

func main() {
	for i := 1; i < 2; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Println(j, j*j, j*j*j, j*j*j*j)

		}

	}
	fmt.Println()
}
