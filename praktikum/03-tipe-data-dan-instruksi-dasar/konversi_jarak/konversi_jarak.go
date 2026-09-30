package main

import "fmt"

func main(){
	var mil float64

	fmt.Println("masukkan jarak: ")
	fmt.Scan(&mil)

	km := mil * 1.6

	fmt.Print(km)
}
