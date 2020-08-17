package main

import (
	"errors"
	"fmt"
)

func main() {
	// sum with array
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)

	//Option ["+", "-", "*", "/"]
	result, err := calculate(0, 0, "=")
	if err != nil {
		fmt.Println("Operasi gagal")
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

func sum(array []int) int {
	result := 0
	for _, score := range array {
		result += score
	}
	return result
}

func calculate(num1, num2 int, operation string) (int, error) {
	var result int
	var err error

	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		err = errors.New("Operasi Tidak ada")
	}

	return result, err

}
