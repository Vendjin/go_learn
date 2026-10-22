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
	operation := inputTypeOperation()
	numbers := inputNumbers()

	switch operation {
	case "AVG":
		fmt.Println(calculateAvg(numbers))
	case "SUM":
		fmt.Println(calculateSum(numbers))
	case "MED":
		fmt.Println(calculateMed(numbers))
	}
}

func inputTypeOperation() string {
	fmt.Println("Выберите тип операции, доступно: AVG, SUM, MED")
	for {
		var typeOperation string
		fmt.Scan(&typeOperation)

		switch typeOperation {
		case "AVG", "SUM", "MED":
			return typeOperation
		default:
			fmt.Println("Не верный тип операции попробуйте снова")
		}
	}

}

func inputNumbers() []int {
	fmt.Println("Введите ваши числа через запятую")
	reader := bufio.NewReader(os.Stdin)

	for {
		numbers, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода, попробуйте снова", err)
			continue
		}

		numbers = strings.TrimSpace(numbers)
		if numbers == "" {
			fmt.Println("Строка пустая, попробуйте снова")
			continue
		}

		arr, err := convertInputToArr(numbers)
		if err != nil {
			fmt.Println(err, "- попробуйте снова")
			continue
		}

		return arr
	}
}

func convertInputToArr(input string) ([]int, error) {
	inputArr := strings.Split(input, ",")
	arr := make([]int, 0, len(inputArr))

	for _, v := range inputArr {
		v = strings.TrimSpace(v)
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("не число: %s", v)
		}

		arr = append(arr, num)
	}

	return arr, nil
}

func calculateSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func calculateAvg(arr []int) float64 {
	return float64(calculateSum(arr)) / float64(len(arr))
}

func calculateMed(arr []int) float64 {
	var med float64

	sort.Ints(arr)
	n := len(arr)

	if n%2 == 1 {
		med = float64(arr[n/2])
	} else {
		med = float64(arr[n/2]+arr[n/2-1]) / 2
	}

	return med
}
