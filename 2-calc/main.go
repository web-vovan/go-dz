package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operation := getOperation()
	numbers, err := getNumbers()

	if err != nil {
		fmt.Println("неверный формат ввода чисел")
	}

	var result float64

	switch operation {
	case "sum":
		result = calculateSum(numbers)
	case "avg":
		result = calculateAvg(numbers)
	case "med":
		result = calculateMed(numbers)
	}

	fmt.Println(result)
}

func calculateSum(numbers []float64) float64 {
	var result float64

	for _, i := range numbers {
		result += i
	}

	return result
}

func calculateAvg(numbers []float64) float64 {
	var total float64

	for _, i := range numbers {
		total += i
	}

	return total / float64(len(numbers))
}

func calculateMed(numbers []float64) float64 {
	sort.Float64s(numbers)

	middle := len(numbers) / 2

	if len(numbers)%2 == 0 {
		return (numbers[middle-1] + numbers[middle]) / 2
	}

	return numbers[middle]
}

func getOperation() string {
	allowOperations := map[string]bool{
		"sum": true,
		"avg": true,
		"med": true,
	}

	var operation string

	for {
		fmt.Print("введите операцию (sum, avg, med): ")
		fmt.Scan(&operation)

		operation = strings.ToLower(operation)

		if _, ok := allowOperations[operation]; !ok {
			fmt.Println("неверная операция")
			continue
		}

		break
	}

	return operation
}

func getNumbers() ([]float64, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("введите числа через запятую (пример: 23, 43, 54): ")

	input, _ := reader.ReadString('\n')
	inputSlice := strings.Split(input, ",")

	numbers := make([]float64, 0, len(inputSlice))

	for _, item := range inputSlice {
		item = strings.TrimSpace(item)

		num, err := strconv.ParseFloat(item, 64)

		if err != nil {
			return numbers, err
		}

		numbers = append(numbers, num)
	}

	return numbers, nil
}
