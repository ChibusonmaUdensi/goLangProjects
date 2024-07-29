package main

import "fmt"

func main() {
	fmt.Println("Enter first citizen name")
	var firstCitizen string
	fmt.Scan(&firstCitizen)
	fmt.Println("Enter first citizen earnings")
	var firstCitizenEarnings float64
	fmt.Scan(&firstCitizenEarnings)

	if firstCitizenEarnings == 30000 {
		var totalTaxForFirstCitizen = firstCitizenEarnings + (0.15 * float64(firstCitizenEarnings))
		fmt.Println("First citizen tax earnings is", totalTaxForFirstCitizen)
	} else if firstCitizenEarnings > 30000 {
		var totalTaxForFirstCitizen = firstCitizenEarnings + (0.20 * float64(firstCitizenEarnings))
		fmt.Println("First citizen tax earnings is", totalTaxForFirstCitizen)
	} else if firstCitizenEarnings < 30000 {
		fmt.Println("First citizen tax earnings is", firstCitizenEarnings)
	}

	fmt.Println("Enter second citizen name")
	var secondCitizen string
	fmt.Scan(&secondCitizen)
	fmt.Println("Enter second citizen earnings")
	var secondCitizenEarnings float64
	fmt.Scan(&secondCitizenEarnings)

	if secondCitizenEarnings == 30000 {
		var totalTaxForSecondCitizen = secondCitizenEarnings + (0.15 * float64(secondCitizenEarnings))
		fmt.Println("Second citizen tax earnings is", totalTaxForSecondCitizen)
	} else if secondCitizenEarnings > 30000 {
		var totalTaxForSecondCitizen = secondCitizenEarnings + (0.20 * float64(secondCitizenEarnings))
		fmt.Println("Second citizen tax earnings is", totalTaxForSecondCitizen)
	} else if secondCitizenEarnings < 30000 {
		fmt.Println("Second citizen tax earnings is", secondCitizenEarnings)
	}
	fmt.Println("Enter third citizen name")
	var thirdCitizen string
	fmt.Scan(&thirdCitizen)
	fmt.Println("Enter third citizen earnings")
	var thirdCitizenEarnings float64
	fmt.Scan(&thirdCitizenEarnings)

	if thirdCitizenEarnings == 30000 {
		var totalTaxForThirdCitizen = thirdCitizenEarnings + (0.15 * float64(thirdCitizenEarnings))
		fmt.Println("Third citizen tax earnings is", totalTaxForThirdCitizen)
	} else if thirdCitizenEarnings > 30000 {
		var totalTaxForThirdCitizen = thirdCitizenEarnings + (0.20 * float64(thirdCitizenEarnings))
		fmt.Println("Third citizen tax earnings is", totalTaxForThirdCitizen)
	} else if thirdCitizenEarnings < 30000 {
		fmt.Println("Third citizen tax earnings is", thirdCitizenEarnings)
	}

}
