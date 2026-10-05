package main

import "fmt"

func main() {

	 var cel float64
	 
	 fmt.Print("Masukan Suhu: ")
	 fmt.Scan(&cel)
	
	 reamur := cel * 4 / 5

	 fmt.Print("Suhu dalam Reamur: ", reamur)
	 
}