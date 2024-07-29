package main

import "fmt"

func main() {
	var total int

	fmt.Println("Enter Number")
	var num int
	fmt.Scan(&num)

	for i := 0; i <= num; i++ {
		fmt.Println("Enter Input")
		var input int
		fmt.Scan(&input)

		total += input
		if total == num {
			fmt.Printf("You have %d to sum %d\n", num, total)
			break
		}
	}

}
