package main

import (
	"fmt"
	"math"
)

func main() {
	var berat float64
	var zona int
	var layanan string

	fmt.Scan(&berat, &zona, &layanan)

	if berat > 30 {
		fmt.Println("Berat melebihi batas maksimum (30 kg)")
		return
	}

	var tarifPerkg int

	switch zona {
	case 1:
		tarifPerkg = 10000
	case 2:
		tarifPerkg = 15000
	case 3:
		tarifPerkg = 25000
	default:
		fmt.Println("Zona tidak valid")
		return
	}

	beratBulat := int(math.Ceil(berat))
	total := beratBulat * tarifPerkg

	if layanan == "Exprss" {
		total = total + total/2
		total += total / 2
	}

	fmt.Println(beratBulat)
	fmt.Println(total)

}
