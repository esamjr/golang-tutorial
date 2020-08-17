package main //change it to main

import (
	"fmt"
)

func main() {
	// // For Standar
	for i := 1; i <= 100; i++ {
		fmt.Println("Saya", i)
	}

	// // While Style
	i := 1
	for i <= 100 {
		fmt.Println("Semangat Belajar!!!", i)
		i++
	}

	//Foreach Style
	title := "Golang the best language"

	for index, letter := range title {
		fmt.Println("index :", index, "letter :", string(letter))
	}

}
