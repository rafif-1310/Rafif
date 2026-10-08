package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	fmt.Scan(&tahun, &bulan)

	// Cek tahun kabisat
	kabisat := (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0

	var hari int

	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		hari = 31
	case "Apr", "Jun", "Sep", "Nov":
		hari = 30
	case "Feb":
		if kabisat {
			hari = 29
		} else {
			hari = 28
		}
	default:
		// Nama bulan tidak valid: tidak ada keluaran ("-")
		fmt.Println("-")
		return
	}

	fmt.Println(hari)
}