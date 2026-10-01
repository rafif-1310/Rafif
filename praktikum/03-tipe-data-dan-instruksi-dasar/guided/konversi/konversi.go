package main

import "fmt"

func main() {
	var cel float64
	fmt.Print("Masukan Suhu: ")
	fmt.Scan(&cel)

	fmt.Print(cel + 273)
}