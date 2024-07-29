package main

import "fmt"

func main() {
	fmt.Println("Enter Account Number")
	var accNo int64
	fmt.Scan(&accNo)
	fmt.Println("Enter beginning balance")
	var beginBalance int64
	fmt.Scan(&beginBalance)
	fmt.Println("Enter total of items charged")
	var totalCharges int64
	fmt.Scan(&totalCharges)
	fmt.Println("Enter total of credits applied")
	var credits int64
	fmt.Scan(&credits)
	fmt.Println("Enter allowed credit limit")
	var allowedCredit int64
	fmt.Scan(&allowedCredit)

	newBalance := beginBalance + totalCharges - credits
	fmt.Println("New balance is", newBalance)
	if newBalance > allowedCredit {
		fmt.Println("Credit Limit Exceeded")
	}
}
