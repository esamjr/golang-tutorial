package unmain //change to main

import "fmt"

func main() {

	// Model Pertama
	var lang [4]string
	lang[0] = "Indonesia"
	lang[1] = "Russia"
	lang[2] = "Korean"
	lang[3] = "Japan"
	fmt.Println(lang)
	length := len(lang)
	fmt.Println(length)

	// Model Kedua
	languages := [5]string{
		"Indonesia",
		"Russia",
		"Korean",
		"Japan",
		"English",
	}
	fmt.Println(languages)
	length = len(languages)
	fmt.Println(length)

	// Model Ketiga
	languages = [...]string{
		"Indonesia",
		"Russia",
		"Korean",
		"Japan",
		"English",
	}
	fmt.Println(languages)
	length = len(languages)
	fmt.Println(length)

	// For each array
	languages = [...]string{
		"Indonesia",
		"Russia",
		"Korean",
		"Japan",
		"English",
	}
	for index, lang := range languages {
		fmt.Println("Index : ", index, " language : ", lang)
	}
}
