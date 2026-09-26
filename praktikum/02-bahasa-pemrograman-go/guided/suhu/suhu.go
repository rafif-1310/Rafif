package main

import "fmt"

func main() {
	var suhu float64
 //Membaca input

 fmt.Print("Masukkan suhu dalam Celcius: ")
 fmt.Scan(&suhu)

 //Menampilkan output
  fmt.Println("Suhu dalam Fahrenheit: ", (suhu*9/5)+32)
  fmt.Println("Suhu dalam Kelvin: ", suhu+273.15)
  fmt.Println("Suhu dalam Reannur: ", suhu*0.8)
  
}
