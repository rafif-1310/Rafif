package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int

	fmt.Scan(&x)

	var hasilKuadrat []string
	totalKuadrat := 0

	for x > 0 {

		digit := x % 10
		kuadrat := digit * digit

		hasilKuadrat = append([]string{strconv.Itoa(kuadrat)}, hasilKuadrat...)

		totalKuadrat += kuadrat

		x = x / 10
	}

	for i, val := range hasilKuadrat {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()

	fmt.Println(totalKuadrat)
}
