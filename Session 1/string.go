package main

import "fmt"

func main() {
	// Tipe data string ditandai dengan nilai yang diapit oleh tanda petik dua (" ")
	var pesan string = "Hallo\nIni baris kedua dengan petik dua.\" <- Ini adalah kutipan"

	// String yang dideklarasikan menggunakan backticks (` `) akan menampilkan semua karakter apa adanya,
	// termasuk karakter spesial seperti \n, tanda kutip dua ("), kutip satu ('), dan baris baru.
	var pesan1 = `Hallo\nIni baris kedua dengan backticks." <- Ini tetap dianggap sebagai bagian dari string`

	// Menampilkan kedua jenis string
	fmt.Println("== Menggunakan tanda petik dua ==")
	fmt.Println(pesan) // karakter \n akan dianggap sebagai baris baru

	fmt.Println("\n== Menggunakan backticks ==")
	fmt.Println(pesan1) // karakter \n tetap ditampilkan apa adanya
}
