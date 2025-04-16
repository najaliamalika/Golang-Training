package main

import "fmt"

func main() {
	// Menggunakan fmt.Println untuk mencetak string dengan newline otomatis di akhir
	fmt.Println("This is my first session in learning Go Programming Language!")

	// Jika ada beberapa argumen dipisahkan dengan koma, fmt.Println akan menambahkan spasi otomatis antar argumen
	fmt.Println("My hobby?", "I love to read!") // untuk membuat space dengan comma

	// fmt.Print tidak menambahkan newline secara otomatis, jadi tambahkan \n di akhir string
	fmt.Print("This is my first session in learning Go Programming Language!\n")

	// Sama seperti sebelumnya, menggunakan fmt.Print dan memisahkan argumen dengan koma
	// tapi perlu tambahkan spasi manual di dalam string agar hasilnya rapi
	fmt.Print("My hobby?", "I love to read!\n") // Perhatikan spasi setelah tanda kutip pertama

	// Menampilkan beberapa string terpisah menggunakan fmt.Print
	// Menambahkan string " " (spasi) di antaranya agar hasil cetakan tidak menyatu
	fmt.Print("Najalia", " ", "Malika\n")
}
