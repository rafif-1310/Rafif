package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	jumlah := 0
	for x > 0 {
		digit := x % 10
		fmt.Print(digit, " ")
		jumlah += digit
		x /= 10
	}
	fmt.Println()
	fmt.Println(jumlah)
}
