package main

import "fmt"

/Project Variable : Membuat Variable With Data Type dan Variable Without Data Type/

func main() {
// Variable With Data Type
var name1 string = ("Najalia") // Mendeklarasikan variabel 'name1' dengan tipe data string dan memberikan nilai "Najalia"
var age int = 21 // Mendeklarasikan variabel 'age' dengan tipe data int dan memberikan nilai 21

// Menampilkan nilai variabel menggunakan fmt.Println
fmt.Println("Ini adalah namanya ==>", name1)
fmt.Println("Ini adalah umurnya ==>", age)


// Variable Without Data Type 

name := "Malika"      // Mendeklarasikan variabel 'name' tanpa menyebutkan tipe data, Go akan otomatis mengenali sebagai string
hobby := "Membaca"    // Mendeklarasikan variabel 'hobby' tanpa tipe data
umur := 20            // Go akan mengenali ini sebagai int

// Menampilkan nilai variabel tanpa tipe data
fmt.Println("Nama tanpa tipe data ==>", name)
fmt.Println("Hobi ==>", hobby)
fmt.Println("Umur ==>", umur)
}


