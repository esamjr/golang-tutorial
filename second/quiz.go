package quiz //change it to main

import (
	"fmt"
)

func main() {

	// Quiz
	title := "Golang the best language"
	// Genap Only
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index :", index, "letter :", string(letter))
		}
	}

	// Vocal word Only
	for index, letter := range title {
		letterString := string(letter)

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("index :", index, "letter :", string(letter))
		}
	}
}
