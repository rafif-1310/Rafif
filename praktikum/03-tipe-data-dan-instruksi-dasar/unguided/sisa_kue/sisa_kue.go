package main

import "fmt"

func main() {

	var x,y int
	
	fmt.Print("Masukkan jumlah kue: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan jumlah anggota keluarga: ")	
	fmt.Scan(&y)

	sisa := x % y

	fmt.Print("Jumlah kue yang tersisa:", sisa )

}