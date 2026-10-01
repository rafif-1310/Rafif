package main

import "fmt"

func main() {
	var p, l int

	fmt.Print("Masukkan panjang (p) dan lebar (l): ")
	 fmt.Scanln(&p, &l)

	// Menghitung luas dan keliling
	luas := p * l
	keliling := 2 * (p + l)

	// Menampilkan hasil keluaran
	fmt.Printf("Keluaran:%d %d\n", luas, keliling)
}