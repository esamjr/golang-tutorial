package main

import (
	"first/calculate"
	"fmt"
)

// Fungsi Main
func main() {
	fmt.Println("Hello, Boys")

	fmt.Println("Input First Number: ")
	var num1 int
	fmt.Scanln(&num1)

	fmt.Println("Enter Second Number: ")
	var num2 int
	fmt.Scanln(&num2)

	result := calculate.Multiple(num1, num2)

	// add := calculate.Add()

	// delete := calculate.Delete(add, 5)

	// result := calculate.Multiple(delete, 5)

	fmt.Println("The Result Is :")
	fmt.Println(result)
}
