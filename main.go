package main

import "fmt"

func main() {
	const usdToEur = 0.87
	const usdToRub = 80.97

	eurToRub := usdToRub / usdToEur

	fmt.Println("1 Eur to RUB", eurToRub)
}
