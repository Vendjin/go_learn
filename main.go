package main

import (
	"fmt"
	"strings"
)

const (
	usdToEur = 0.87
	usdToRub = 80.97
)

func main() {
	fmt.Println("===Конвертер валют===")
	baseCurrency, ok := getBaseCurrency()
	if !ok {
		fmt.Println("Завершение работы нв выборе базовой валюты")
		return
	}

	amount, ok := getAmount()
	if !ok {
		fmt.Println("Завершение работы на вводе суммы")
		return
	}

	targetCurrency, ok := getTargetCurrency(baseCurrency)
	if !ok {
		fmt.Println("Завершение работы на выборе конечной валюты.")
		return
	}

	result := calculateCurrency(amount, baseCurrency, targetCurrency)
	fmt.Printf("\nРезультат: %.2f %s = %.2f %s\n", amount, baseCurrency, result, targetCurrency)

}

func getBaseCurrency() (string, bool) {
	for {
		fmt.Println("Выберите вашу валюту: RUB, EUR, USD")
		var currency string
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)

		if !checkInputCurrency(currency) {
			fmt.Printf("Недопустимая валюта: \"%s\".\n", currency)
		} else {
			return currency, true
		}

		if !checkRepeatChoice() {
			return "", false
		}
	}
}

func getTargetCurrency(baseCurrency string) (string, bool) {
	for {
		fmt.Printf("Выберите целевую валюту из RUB, EUR, USD, кроме %v:", baseCurrency)
		var currency string
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)

		if !checkInputCurrency(currency) {
			fmt.Printf("Недопустимая валюта: \"%s\".\n", currency)
		} else if currency == baseCurrency {
			fmt.Printf("Валюта перевода не может совпадать с исходной валютой %s!\n", baseCurrency)
		} else {
			return currency, true
		}

		if !checkRepeatChoice() {
			return "", false
		}
	}
}

func getAmount() (float64, bool) {
	for {
		fmt.Print("Введите сумму для конвертации: ")
		var amount float64
		_, err := fmt.Scan(&amount)

		if err == nil && amount > 0 {
			return amount, true
		}

		fmt.Println("Ошибка: Введите положительное число!")

		var discard string
		fmt.Scanln(&discard)

		if !checkRepeatChoice() {
			return 0, false
		}
	}
}

func checkInputCurrency(val string) bool {
	switch strings.ToLower(val) {
	case "rub", "eur", "usd":
		return true
	default:
		return false
	}
}

func checkRepeatChoice() bool {
	var userChoice string
	fmt.Println("Вы хотите повторить ввод? (y/n):")
	fmt.Scan(&userChoice)

	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}

func calculateCurrency(amount float64, srcCurrency string, targetCurrency string) float64 {
	var amountInUSD float64

	switch srcCurrency {
	case "USD":
		amountInUSD = amount
	case "EUR":
		amountInUSD = amount / usdToEur
	case "RUB":
		amountInUSD = amount / usdToRub
	}

	switch targetCurrency {
	case "USD":
		return amountInUSD
	case "EUR":
		return amountInUSD * usdToEur
	case "RUB":
		return amountInUSD * usdToRub
	default:
		return 0
	}
}
