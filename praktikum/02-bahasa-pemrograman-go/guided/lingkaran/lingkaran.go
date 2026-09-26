package main

import "fmt"

func main() {
	var jari float64
	var pi = 3.14
	var luas float64

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&jari)

	luas = pi * jari * jari

    fmt.Println("================ OUTPUT ==================")
	fmt.Println("Jari-jari           :", jari)
	fmt.Println("Luas lingkaran      :", luas)
	fmt.Println("==========================================")
}