package fourth //Change To (main) to use

func main() {
	//Fungsi Result
	// sentence := printResult("Saya")
	// fmt.Println(sentence)

	//Fungsi add
	// var number1 int
	// fmt.Println("Choose First Number : ")
	// fmt.Scanln(&number1)
	// var number2 int
	// fmt.Println("Choose Second Number : ")
	// fmt.Scanln(&number2)
	// result := add(number1, number2)
	// fmt.Println("The result Is :")
	// fmt.Println(result)

	//Fungsi calculate
	// luas, _ := calculate(10, 2)
	// fmt.Println(luas)
	// // fmt.Println(keliling)

}

// Fungsi printResult
func printResult(sentence string) string {
	newSentence := sentence + " Bermain golang"
	return newSentence
}

// Fungsi Add
func add(number, numberTwo int) int {
	result := number + numberTwo
	return result
}

// Fungsi calculate
func calculate(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

// Predefine return value
func PreCalculate(panjang, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	return luas, keliling
}
