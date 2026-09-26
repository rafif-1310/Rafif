 package main
 import "fmt"

 func main(){ 
	var nama string
	var skorMatematika, skorBahasaInggris int
	
	//Membaca input
	fmt.Scan(&nama)
	fmt.Scan(&skorMatematika)
	fmt.Scan(&skorBahasaInggris)

	//Menghitung total & rata-rata skor (Pembagian bilangan bulat) 
	total := skorMatematika + skorBahasaInggris
	rataRata := total / 2

	//Menampilkan output
	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(rataRata)
	
 }