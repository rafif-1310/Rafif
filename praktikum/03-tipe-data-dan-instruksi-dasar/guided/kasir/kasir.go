package main

import "fmt"

func main() {
	var uang int
	fmt.Print("Uang = ")
	fmt.Scan(&uang)

	sepuluh := uang / 10000
	sisa := uang % 10000
	lima := sisa / 5000
	sisa = sisa % 5000
	satu := sisa / 1000
	sisa = sisa % 1000
	fmt.Println("10000 = ", sepuluh)
	fmt.Println("5000 = ", lima)
	fmt.Println("1000 = ", satu)
	fmt.Print("sisa = ", sisa)
}