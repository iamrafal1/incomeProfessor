package main

import "fmt"

var STD_PAYE_RATE float32 = 36800

func main() {
	fmt.Println("Hello World")
	fmt.Println(yearlyPayeCaluclator(65000))
}

func yearlyPayeCaluclator(grossPay float32) float32 {
	if grossPay < STD_PAYE_RATE {
		return grossPay * 0.8
	}
	higherBracketTax := (grossPay - STD_PAYE_RATE) * 0.4
	lowerBracketTax := STD_PAYE_RATE * 0.2

	fmt.Println(higherBracketTax, lowerBracketTax)

	return grossPay - (higherBracketTax + lowerBracketTax)
}
