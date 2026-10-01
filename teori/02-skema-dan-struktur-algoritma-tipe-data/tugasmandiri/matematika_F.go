package main

import "fmt"

func main() {
	var x float64

	// Meminta masukan dari pengguna
	fmt.Print("Masukkan sebuah bilangan berkoma x: ")
	_, err := fmt.Scanln(&x)
	if err != nil {
		fmt.Println("Masukan tidak valid")
		return
	}

	// Memeriksa jika pembagi adalah nol
	if x == 3 {
		fmt.Println("Error: Nilai x tidak boleh 3 karena menyebabkan pembagian dengan nol.")
		return
	}

	// Menghitung f(x)
	hasil := (x*x + 2*x + 1) / (x - 3)

	// hasil keluaran
	fmt.Printf("Keluaran: %v\n", hasil)
}