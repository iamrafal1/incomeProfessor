package main

import "fmt"

var StdPayeRate float32 = 36800
var TaxCredits float32 = 3400

func main() {
	fmt.Println("Hello World")
	var myPay float32 = 48360
	var paye float32 = yearlyPayeCaluclator(myPay)
	var usc float32 = yearlyUscCaluclator(myPay)
	var prsi float32 = yearlyPrsiCalculator(myPay)
	fmt.Println("PAYE: ", paye)
	fmt.Println("USC: ", usc)
	fmt.Println("PRSI: ", prsi)
	fmt.Println("net pay: ", myPay-(usc+paye+prsi))
}

func yearlyPayeCaluclator(grossPay float32) float32 {
	if grossPay < StdPayeRate {
		return (grossPay * 0.2) - TaxCredits
	}
	higherBracketTax := (grossPay - StdPayeRate) * 0.4
	lowerBracketTax := StdPayeRate * 0.2

	fmt.Println(higherBracketTax, lowerBracketTax)

	return higherBracketTax + lowerBracketTax - TaxCredits
}

func yearlyUscCaluclator(grossPay float32) float32 {
	if grossPay < 13000 {
		return grossPay
	}
	var halfPercent float32 = 12012 * 0.005
	if grossPay < 21295 {
		return ((grossPay - 12012) * 0.02) + halfPercent
	}
	var twoPercent float32 = (21295 - 12012) * 0.02
	if grossPay < 70044 {
		return ((grossPay - 21295) * 0.045) + halfPercent + twoPercent
	}
	var fourAndHalfPercent float32 = (70044 - 21295) * 0.045
	return ((grossPay - 70044) * 0.08) + halfPercent + twoPercent + fourAndHalfPercent
}

func yearlyPrsiCalculator(grossPay float32) float32 {
	if grossPay < 18304 {
		return 0
	}
	return grossPay * 0.04
}
