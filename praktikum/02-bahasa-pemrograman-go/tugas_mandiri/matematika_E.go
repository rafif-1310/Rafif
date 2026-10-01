package main

import "fmt"

func main() {
	var x, y int

	// Membaca input x dan y
	fmt.Scan(&x, &y)

	// Konversi ke float64 untuk perhitungan pecahan
	xf := float64(x)
	yf := float64(y)

	// Menghitung f(x, y) = 5x^2 - 2xy + (y^3 / (x + 1))
	hasil := 5*(xf*xf) - 2*xf*yf + (yf*yf*yf)/(xf+1)

	// Mencetak hasil dengan format desimal
	fmt.Println(hasil)
}

