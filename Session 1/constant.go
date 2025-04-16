package main

import "fmt"

func main() {
	// Mendeklarasikan konstanta dengan tipe data string
	// const artinya nilai variabel ini tidak bisa diubah setelah ditetapkan
	const full_Name string = "Najalia Malika"

	// Menggunakan fmt.Printf untuk mencetak teks dengan format
	// %s adalah format verb untuk menampilkan nilai string
	fmt.Printf("Hello %s", full_Name)
}
