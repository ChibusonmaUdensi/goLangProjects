package main

import "fmt"

func main() {
	var totalItems float64 = 0
	var add string
	for {
		fmt.Println("Add more items? (yes/no)")

		fmt.Scan(&add)
		if add != "yes" {
			break
		}
		var item float64
		fmt.Scan(&item)
		totalItems += item

	}
	fmt.Println(totalItems)
	var commission = 200 + (0.09 * float64(totalItems))
	fmt.Println("Your Sales Commission is", commission)

}
