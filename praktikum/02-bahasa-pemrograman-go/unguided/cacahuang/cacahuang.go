package main

import "fmt"

func main() {

var Jumlah int32
fmt.Println("Masukkan jumlah uang: ")
fmt.Print("jumlah uang ")
fmt.Scan(&Jumlah)

uang10000 := Jumlah / 10000
sisa := Jumlah % 10000
uang5000 := sisa / 5000
sisa = sisa % 5000
uang1000 := sisa / 1000
sisa = sisa % 1000

fmt.Println("Jumlah uang:", Jumlah)
fmt.Println("Uang 10000:", uang10000)
fmt.Println("Uang 5000 :", uang5000)
fmt.Println("Uang 1000 :", uang1000)
fmt.Println("Sisa uang  :", sisa)
}