package main

import "fmt"

func main() {
	var n, lebar int
	var tersempit int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&lebar)

		if i == 1 {
			tersempit = lebar
		} else if lebar < tersempit {
			tersempit = lebar
		}
	}

	fmt.Println("Tersempit :", tersempit)
}