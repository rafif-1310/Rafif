package main

import (
	"fmt"
)

func main() {
	var K int

	fmt.Scan(&K)

	totalBarang := 0

	for i := 0; i < 5; i++ {
		var V int
		fmt.Scan(&V)

		totalBarang += V

		if totalBarang >= K {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	}
}
