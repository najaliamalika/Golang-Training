package main

import "fmt"

func main() {
	var benar = true
	var salah = false

	// Operator AND (&&): akan bernilai true hanya jika kedua kondisi bernilai true
	var hasilAND = benar && benar
	fmt.Printf("benar && benar \t= %t\n", hasilAND)

	// Operator OR (||): akan bernilai true jika salah satu kondisi bernilai true
	var hasilOR = salah || benar
	fmt.Printf("salah || benar \t= %t\n", hasilOR)

	// Operator NOT (!): membalik nilai boolean (true jadi false, false jadi true)
	var hasilNOT = !salah
	fmt.Printf("!salah \t\t= %t\n", hasilNOT)
}
