
package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)

	ascending := true
	descending := true

	temp := angka

	for temp >= 10 {
		digit1 := temp % 10
		digit2 := (temp / 10) % 10

		if digit2 > digit1 {
			ascending = false
		}

		if digit2 < digit1 {
			descending = false
		}

		temp = temp / 10
	}

	if ascending {
		fmt.Println("ascending")
	} else if descending {
		fmt.Println("descending")
	} else {
		fmt.Println("ngotak dikit kontol ga masuk apa apa memek")
	}
}
