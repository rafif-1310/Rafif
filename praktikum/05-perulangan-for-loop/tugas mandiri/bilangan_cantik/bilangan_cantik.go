package main

import (
	"fmt"
)

func main() {
	var X int
	fmt.Print("Masukkan bilangan : ")
	_, err := fmt.Scan(&X)
	if err != nil || X <= 0 {
		fmt.Println("Input harus berupa bilangan bulat positif!")
		return
	}

	isCantik := true
	lastDigit := 10

	for X > 0 && isCantik {
		currentDigit := X % 10

		if currentDigit > lastDigit {
			isCantik = false
		} else {
			lastDigit = currentDigit
		}
		X = X / 10
	}

	fmt.Println("Keluaran:", isCantik)
}
