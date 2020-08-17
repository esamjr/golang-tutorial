package ifelse //change it to main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("Input Your Age: ")
	fmt.Scanln(&age)

	var status string

	if age > 10 {
		status = "Udah Puber"
	} else if age <= 10 {
		status = "Masih Kecil dek"
	} else {
		status = "Isi Umur kamu!!!"
	}

	fmt.Println(status)
}
