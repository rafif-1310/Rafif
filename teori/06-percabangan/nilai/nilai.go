package main

import "fmt"

func main() {
	var nilai int
	fmt.Scan(&nilai)

	if nilai >= 85 {
		fmt.Println("A")
	} else if nilai >= 70 {
		fmt.Println("B")
	} else if nilai >= 50 {
		fmt.Println("C")
	} else if nilai < 50 {
		fmt.Println("E")
	}
}
