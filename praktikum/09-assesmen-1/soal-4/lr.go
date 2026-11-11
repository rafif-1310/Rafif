package main

import "fmt"

func prima(angka int) bool {
	if angka < 2 {
		return false
	}

	for i := 2; i < angka; i++ {
		if angka%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var L, R int
	var jumlah int

	fmt.Scan(&L, &R)

	for angka := L; angka <= R; angka++ {
		temp := angka
		totalDigit := 0
		sebelumnya := 10
		valid := true

		if temp == 0 {
			totalDigit = 0
		}

		for temp > 0 {
			digit := temp % 10

			totalDigit = totalDigit + digit

			if digit > sebelumnya {
				valid = false
			}

			sebelumnya = digit
			temp = temp / 10
		}

		if valid && prima(totalDigit) {
			jumlah++
		}
	}

	fmt.Println(jumlah)
}