package main

import (
	"fmt"
	"strconv"
)

const usdToRub float64 = 80.75
const eurToRub float64 = 94.46

func main() {
	convertCurrency()
}

func setCurrency() string {
	var currency string
	for {
		fmt.Println("Введите валюту - usd, eur, rub: ")
		fmt.Scanln(&currency)
		if currency == "usd" || currency == "eur" || currency == "rub" {
			return currency
		}
		fmt.Println("❌ Ошибка: неизвестная валюта, попробуйте снова.")
	}
}

func setConvertCurrency() string {
	var convertCurrency string
	for {
		fmt.Println("Введите валюту в которую нужно перевести - usd, eur, rub: ")
		fmt.Scanln(&convertCurrency)
		if convertCurrency == "usd" || convertCurrency == "eur" || convertCurrency == "rub" {
			return convertCurrency
		}
		fmt.Println("❌ Ошибка: неизвестная валюта, попробуйте снова.")
	}
}

func setCurrencyValue(currency string) float64 {
	var input string
	for {
		fmt.Println("Введите значение в " + currency + ", которое хотите конвертировать: ")
		fmt.Scanln(&input)
		value, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return value
		}
		fmt.Println("❌ Ошибка: нужно ввести число, попробуйте снова.")
	}
}

func convertCurrency() {
	var rubCurrencyValue float64
	var convertValue float64

	currency := setCurrency()
	currencyValue := setCurrencyValue(currency)
	targetCurrency := setConvertCurrency()

	// переводим всё в рубли
	switch currency {
	case "rub":
		rubCurrencyValue = currencyValue
	case "eur":
		rubCurrencyValue = currencyValue * eurToRub
	case "usd":
		rubCurrencyValue = currencyValue * usdToRub
	default:
		rubCurrencyValue = 0
	}

	// конвертируем из рублей в целевую валюту
	switch targetCurrency {
	case "rub":
		convertValue = rubCurrencyValue
	case "eur":
		convertValue = rubCurrencyValue / eurToRub
	case "usd":
		convertValue = rubCurrencyValue / usdToRub
	default:
		convertValue = 0
	}

	// вывод результата
	fmt.Printf("✅ %.2f %s = %.2f %s\n", currencyValue, currency, convertValue, targetCurrency)
}
