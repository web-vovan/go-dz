package main

import (
	"bufio"
	"errors"
    "strconv"
    "strings"
	"fmt"
	"os"
)

const USD_EUR = 0.91
const USD_RUB = 84.24

func main() {
	var originalCurrency string
	var targetingCurrency string
	var quantity float64
	var result float64
	var err error

	originalCurrency = getOriginalCurrency()
	quantity = getQuantity()
	targetingCurrency = getTargetCurrency(originalCurrency)

	result, err = calculate(originalCurrency, targetingCurrency, quantity)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%.2f %s = %.2f %s", quantity, originalCurrency, result, targetingCurrency)
	}
}

func calculate(originalCurrency string, targetingCurrency string, quantity float64) (float64, error) {
	if originalCurrency == "usd" {
		switch targetingCurrency {
		case "rub":
			return USD_RUB * quantity, nil
		case "eur":
			return USD_EUR * quantity, nil
		}
	}

	if originalCurrency == "eur" {
		switch targetingCurrency {
		case "usd":
			return quantity / USD_EUR, nil
		case "rub":
			return (USD_RUB / USD_EUR) * quantity, nil
		}
	}

	if originalCurrency == "rub" {
		switch targetingCurrency {
		case "usd":
			return quantity / USD_RUB, nil
		case "eur":
			return quantity / (USD_RUB / USD_EUR), nil
		}
	}

	return 0, errors.New("ошибка в расчетах")
}

func getOriginalCurrency() string {
	var originalCurrency string

	for {
		fmt.Print("исходная валюта(usd, eur, rub): ")
		fmt.Scan(&originalCurrency)

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
