package unmain //change it to main

import "fmt"

func main() {
	// Method 1
	var myMap map[string]int
	myMap = map[string]int{}

	myMap["jaksel"] = 255
	myMap["Jakbar"] = 256
	myMap["jaktim"] = 265

	fmt.Println(myMap)

	// method 2
	myMap1 := map[string]string{
		"jaksel": "Selatan",
		"jakbar": "Barat",
		"Jaktim": "Timur",
	}
	fmt.Println(myMap1)

	// For range map
	myMap2 := map[string]string{
		"jaksel": "Selatan",
		"jakbar": "Barat",
		"Jaktim": "Timur",
	}
	for key, value := range myMap2 {
		fmt.Println("Key : ", key, " value: ", value)
	}

	fmt.Println("===========")

	// Delete Map
	delete(myMap2, "jaksel")
	for key, value := range myMap2 {
		fmt.Println("Key : ", key, " value: ", value)
	}

	value, isAvailabe := myMap2["jakbar "]
	fmt.Println(isAvailabe)
	fmt.Println(value)
}
