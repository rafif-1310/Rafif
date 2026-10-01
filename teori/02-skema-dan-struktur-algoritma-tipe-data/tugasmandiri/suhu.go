package main

import "fmt"

func main() {
	var c float64

	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scanln(&c)
	
	// Menghitung konversi ke Fahrenheit
	f := c * 9 / 5 + 32

	// Menampilkan hasil dengan 1 angka di belakang koma
	fmt.Printf("Keluaran: %.1f\n", f)
}