package main

import "fmt"

func main() {
	var totalHari int
	fmt.Print("Masukkan jumlah hari: ")
	fmt.Scanln(&totalHari)

	tahun := totalHari / 360
	bulan := (totalHari % 360) / 30
	minggu := ((totalHari % 360) % 30) / 7
	hari := ((totalHari % 360) % 30) % 7

	fmt.Printf("%d\n%d\n%d\n%d\n", tahun, bulan, minggu, hari)
}