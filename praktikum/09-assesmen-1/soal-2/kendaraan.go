package main

import "fmt"

func main() {
	var kendaraan string
	var motor, mobil, truk int

	for {
		fmt.Scan(&kendaraan)

		if kendaraan == "M" {
			motor++
		} else if kendaraan == "K" {
			mobil++
		} else if kendaraan == "T" {
			truk++
		} else {
			break
		}
	}

	fmt.Println("Tipe M =", motor)
	fmt.Println("Tipe K =", mobil)
	fmt.Println("Tipe T =", truk)
}
