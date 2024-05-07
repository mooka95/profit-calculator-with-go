package main

import "fmt"

func main() {
	var revene float64
	var taxRate float64
	var expenses float64
	fmt.Println("Please enter revene")
	fmt.Scan(&revene)
	fmt.Println("Please enter Tax Rate with %")
	fmt.Scan(&taxRate)
	fmt.Println("Please enter expenses")
	fmt.Scan(&expenses)
	earningBeforeTax := revene - expenses
	earningAfterTax := earningBeforeTax - ((taxRate / 100) * earningBeforeTax)
	ratio := earningBeforeTax / earningAfterTax
	fmt.Printf("earing before tax is %v AND earning after tax is %v ratio is %v", earningBeforeTax, earningAfterTax, ratio)
}
