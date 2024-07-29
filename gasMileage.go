package main

import "fmt"

func main() {

	fmt.Println("Enter the miles driven")
	var miles int
	fmt.Scanf("%d", &miles)

	fmt.Println("Enter the gallon used")
	var gallon int
	fmt.Scanf("%d", &gallon)

	mpg := float64(miles) / float64(gallon)

	fmt.Println("The miles per gallon is", mpg)
}
