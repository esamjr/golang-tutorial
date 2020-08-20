package main

import "fmt"

// Interface
type BangunDatar interface {
	HitungLuas() int
}

// Struct 1
type Persegi struct {
	Sisi int
}

// Method Untuk Interface
func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

// Struct 2
type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

// Method Untuk Interface
func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	persegi := Persegi{Sisi: 4}
	persegiPanjang := PersegiPanjang{Panjang: 5, Lebar: 1000}

	luas := SeberapaLuas(persegi)
	fmt.Println("Luas Persegi : ", luas)

	luas = SeberapaLuas(persegiPanjang)
	fmt.Println("Luas Persegi Panjang : ", luas)

}

// Function Dengan Interface
func SeberapaLuas(bangunDatar BangunDatar) int { //Menggunakan Interface Luas
	return bangunDatar.HitungLuas()
}
