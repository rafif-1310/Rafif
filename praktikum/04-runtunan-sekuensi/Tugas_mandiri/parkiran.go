package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Parkir Kampus")
	
	for kendaraan := 1; kendaraan <= 50; kendaraan++ {
		
		fmt.Printf("Kendaraan datang-%d\n", kendaraan)
		
	
		time.Sleep(50 * time.Millisecond) 
	}

	
	fmt.Println("\nMaksimum 50")
	fmt.Println("Parkir Penuh")
}