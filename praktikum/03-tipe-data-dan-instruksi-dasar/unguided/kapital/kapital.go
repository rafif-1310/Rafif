package main
import ("fmt")

func main() {
	var huruf rune
	fmt.Print("Masukkan sebuah huruf: ")
	fmt.Scanf("%c", &huruf)

	if huruf >= 'A' && huruf <= 'Z' {
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}
}