package unmain //Change it To main

import "fmt"

func main() {
	students := []map[string]string{
		{"name": "John", "score": "A"},
		{"name": "Imhoteb", "score": "C"},
		{"name": "Rudolf", "score": "S"},
	}

	fmt.Println(students)
	for _, siswa := range students {
		fmt.Println(siswa["name"], " - ", siswa["score"])
	}
}
