package main

import "fmt"

const USDToEUR = 0.8779
const USDToRUB = 84.34
const EURToRUB = (1 / USDToEUR) * USDToRUB

func main() {

}

func getExchangeData() (string, string) {
	var currency string
	var exchangeCurency string

	fmt.Print("Введите валюту: ")
	fmt.Scan(&currency)

	fmt.Print("Введите валюту перевода: ")
	fmt.Scan(&exchangeCurency)

	return currency, exchangeCurency
}

func convertCurrency(course float64, currency, exchangeCurency string) {

}
