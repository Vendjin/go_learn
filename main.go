package main

import "fmt"

func main() {
	const usdToEur = 0.87
	const usdToRub = 80.97

	eurToRub := usdToRub / usdToEur

	fmt.Println("1 Eur to RUB", eurToRub)
}

func userInput() string {
	var input string
	fmt.Scan(&input)

	return input
}

func calculateCurrency(userNumber string, srcCurrency float64, targetCurrency float64) {}
