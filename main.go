package main

import "fmt"

const USDToEUR = 0.8779
const USDToRUB = 84.34
const EURToRUB = (1 / USDToEUR) * USDToRUB

func main() {
	sum, currency, exchangeCurency := getExchangeData()

	convertCurrency(sum, currency, exchangeCurency)
}

func getExchangeData() (sum float64, currency, exchangeCurency string) {
	fmt.Print("Введите валюту: ")
	fmt.Scan(&currency)

	fmt.Print("Введите валюту перевода: ")
	fmt.Scan(&exchangeCurency)

	fmt.Print("Введите сумму перевода: ")
	fmt.Scan(&sum)

	return
}

func convertCurrency(sum float64, currency, exchangeCurency string) {

}
