package main

import "fmt"

func main() {
	var decimalNumber float32 = 3.63

	// verb %f digunakan dalam fmt.Printf untuk memformat angka desimal (float)
	fmt.Printf("Decimal Number (default float): %f\n", decimalNumber)

	// Menampilkan hanya 3 digit di belakang koma
	fmt.Printf("Decimal Number (3 digits): %.3f\n", decimalNumber)

	// Menampilkan hanya 1 digit di belakang koma
	fmt.Printf("Decimal Number (1 digit): %.1f\n", decimalNumber)

	// Menampilkan 0 digit di belakang koma (dibulatkan)
	fmt.Printf("Decimal Number (no decimal): %.0f\n", decimalNumber)

	// Menambahkan lebar minimum tampilan total karakter (misalnya: 10 karakter)
	fmt.Printf("Decimal Number (width 10, 2 digits): %10.2f\n", decimalNumber)
}
