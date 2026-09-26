package main

import "fmt"

func main() {

	var a, b int


	fmt.Scan(&a)
	fmt.Scan(&b)


	fmt.Println("Hasil penjumlahan:", a + b)
	fmt.Println("Hasil pengurangan:", a - b)
	fmt.Println("Hasil perkalian:", a * b)
	fmt.Println("Hasil pembagian:", a / b)
	fmt.Println("Hasil sisa bagi:", a % b)

}