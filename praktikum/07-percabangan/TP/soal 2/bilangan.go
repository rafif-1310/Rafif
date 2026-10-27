package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("0 adalah bilangan 0")
	} else {
		if n %2 == 0 {
			fmt.Print("genap, ")
		} else {
			fmt.Print("ganjil, ")
		}
	
		if n > 0 {
			fmt.Print("positif")
		} else {
			fmt.Print("negatif")
		}
	
}
}