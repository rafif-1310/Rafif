package main

import "fmt"

func main() {
	var umur int8
	var suhu float32

	//suhu = 36.3
	//umur = 10
	fmt.Scan(&umur)
	fmt.Scan(&suhu)

	fmt.println("Umur: ", umur)
	fmt.println("Suhu: ", suhu)
	fmt.Println("Alamat memori dari Var suhu ", suhu)
	fmt.Println("Alamat memori dari Var umur ", umur)
}