package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var operations = map[string]func([]float64) float64{
		"sum": calculateSum,
		"avg": calculateAvg,
		"med": calculateMed,
	}

	operation := getOperation()
	numbers, err := getNumbers()

	if err != nil {
		fmt.Println("неверный формат ввода чисел")
	}

	calculateFunc := operations[operation]
	result := calculateFunc(numbers)

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

	input, err := reader.ReadString('\n')

	if err != nil {
		return nil, errors.New("не удалось прочитать строку")
	}

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
