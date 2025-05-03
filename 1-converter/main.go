package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mapCurrencyRate = map[string]map[string]float64

const USD_EUR = 0.91
const USD_RUB = 84.24

func main() {
	var originalCurrency string
	var targetingCurrency string
	var quantity float64
	var result float64
	var err error

	currencyRate := getCurrencyRate()

	originalCurrency = getOriginalCurrency()
	quantity = getQuantity()
	targetingCurrency = getTargetCurrency(originalCurrency)

	result, err = calculate(&currencyRate, originalCurrency, targetingCurrency, quantity)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%.2f %s = %.2f %s", quantity, originalCurrency, result, targetingCurrency)
	}
}

func getCurrencyRate() mapCurrencyRate {
	result := mapCurrencyRate{}

	result["usd"] = map[string]float64{
		"eur": USD_EUR,
		"rub": USD_RUB,
	}

	result["eur"] = map[string]float64{
		"usd": 1 / USD_EUR,
		"rub": USD_RUB / USD_EUR,
	}

	result["rub"] = map[string]float64{
		"usd": 1 / USD_RUB,
		"eur": 1 / (USD_RUB / USD_EUR),
	}

	return result
}

func calculate(currencyRate *mapCurrencyRate, originalCurrency string, targetingCurrency string, quantity float64) (float64, error) {
	rate := (*currencyRate)[originalCurrency][targetingCurrency]

	if rate == 0 {
		return 0, fmt.Errorf("не найден курс %s к %s", originalCurrency, targetingCurrency)
	}

	return rate * quantity, nil
}

func getOriginalCurrency() string {
	var originalCurrency string

	for {
		fmt.Print("исходная валюта(usd, eur, rub): ")
		fmt.Scan(&originalCurrency)

		originalCurrency = strings.ToLower(originalCurrency)

		if originalCurrency == "usd" || originalCurrency == "eur" || originalCurrency == "rub" {
			return originalCurrency
		}

		fmt.Println("неверный тип валюты")
	}
}

func getTargetCurrency(originalCurrency string) string {
	var targetCurrency string

	for {
		fmt.Print("конечная валюта(usd, eur, rub): ")
		fmt.Scan(&targetCurrency)

		targetCurrency = strings.ToLower(targetCurrency)

		if targetCurrency == originalCurrency {
			fmt.Println("валюта не может равняться исходной")
			continue
		}

		if targetCurrency == "usd" || targetCurrency == "eur" || targetCurrency == "rub" {
			return targetCurrency
		}

		fmt.Println("неверный тип валюты")
	}
}

func getQuantity() float64 {
	var quantity float64
	var err error

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("сумма: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		quantity, err = strconv.ParseFloat(input, 64)

		if err == nil {
			return quantity
		}

		fmt.Println("некорректная сумма")
	}
}
