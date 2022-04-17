package main

import "fmt"

var STD_PAYE_RATE float32 = 36800

func main() {
	fmt.Println("Hello World")
	var myPay float32 = 65000
	var paye float32 = yearlyPayeCaluclator(myPay)
	var usc float32 = yearlyUscCaluclator(myPay)
	fmt.Println("PAYE: ", paye)
	fmt.Println("USC: ", usc)
	fmt.Println("net pay: ", myPay-(usc+paye))
}

func yearlyPayeCaluclator(grossPay float32) float32 {
	if grossPay < STD_PAYE_RATE {
		return grossPay * 0.8
	}
	higherBracketTax := (grossPay - STD_PAYE_RATE) * 0.4
	lowerBracketTax := STD_PAYE_RATE * 0.2

	fmt.Println(higherBracketTax, lowerBracketTax)

	return higherBracketTax + lowerBracketTax
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
