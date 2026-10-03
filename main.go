package main

import (
	"fmt"
	"strconv"
)

const USDToEUR = 0.8779
const USDToRUB = 84.34
const EURToRUB = (1 / USDToEUR) * USDToRUB

const USD = "USD"
const EUR = "EUR"
const RUB = "RUB"

func main() {
	sum, currency, exchangeCurency := getExchangeData()

	convertCurrency(sum, currency, exchangeCurency)
}

func getExchangeData() (count float64, currency, exchangeCurency string) {
	currency = inputCurrency()

	count = inputCount()

	exchangeCurency = inputExchangeCurrency(currency)

	return
}

func inputCurrency() string {
	var currency string

	for {
		fmt.Print("Введите валюту ")
		showCurrency("")
		fmt.Scan(&currency)

		isGetCurrency := checkUserCurrency(currency)

		if isGetCurrency {
			return currency
		} else {
			fmt.Println()
			fmt.Println("Неверная валюта, попробуйте снова")
			fmt.Println()
			continue
		}
	}
}

func inputExchangeCurrency(currency string) string {
	var exchangeCurency string

	for {
		fmt.Print("Введите валюту перевода ")
		showCurrency(currency)
		fmt.Scan(&exchangeCurency)

		isGetExchange := checkUserExchangeCurrency(currency, exchangeCurency)

		if isGetExchange {
			return exchangeCurency
		} else {
			fmt.Println()
			fmt.Println("Неверная валюта, попробуйте снова")
			fmt.Println()
			continue
		}
	}
}

func inputCount() float64 {
	var sumStr string
	var sum float64

	for {
		fmt.Print("Введите сумму перевода: ")
		fmt.Scan(&sumStr)

		num, err := strconv.ParseFloat(sumStr, 64)
		if err != nil {
			continue
		} else if num <= 0 {
			fmt.Println()
			fmt.Println("Сумма должна быть больше 0")
			fmt.Println()
		} else {
			sum = num
			return sum
		}

	}
}

func convertCurrency(count float64, currency, exchangeCurrency string) {
	var sum float64

	switch currency {
	case USD:
		if exchangeCurrency == EUR {
			sum = count * USDToEUR
		} else if exchangeCurrency == RUB {
			sum = count * USDToRUB
		}
	case EUR:
		if exchangeCurrency == USD {
			sum = count / USDToEUR
		} else if exchangeCurrency == RUB {
			sum = count * EURToRUB
		}
	case RUB:
		if exchangeCurrency == USD {
			sum = count / USDToRUB
		} else if exchangeCurrency == EUR {
			sum = count / EURToRUB
		}
	}

	fmt.Println()
	fmt.Printf("Итог: %.2f\n", sum)
}

func checkUserCurrency(nameCurrency string) bool {
	switch nameCurrency {
	case USD, EUR, RUB:
		return true
	default:
		return false
	}
}

func checkUserExchangeCurrency(nameCurrency, nameExchangeCurrency string) bool {
	switch nameCurrency {
	case USD:
		if nameExchangeCurrency == EUR || nameExchangeCurrency == RUB {
			return true
		} else {
			return false
		}
	case EUR:
		if nameExchangeCurrency == USD || nameExchangeCurrency == RUB {
			return true
		} else {
			return false
		}
	case RUB:
		if nameExchangeCurrency == USD || nameExchangeCurrency == EUR {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

func showCurrency(nameCurrency string) {
	switch nameCurrency {
	case USD:
		output := fmt.Sprintf("(возможные валюты): %s %s", EUR, RUB)
		fmt.Println(output)
	case EUR:
		output := fmt.Sprintf("(возможные валюты): %s %s", USD, RUB)
		fmt.Println(output)
	case RUB:
		output := fmt.Sprintf("(возможные валюты): %s %s", USD, EUR)
		fmt.Println(output)
	default:
		output := fmt.Sprintf("(возможные валюты): %s %s %s", USD, EUR, RUB)
		fmt.Println(output)
	}
}
