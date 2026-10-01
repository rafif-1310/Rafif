package main
import "fmt"

func main() {
	// kamus
	var x, y int
	var f float64

	//algoritma
	fmt.Scan(&x, &y)

	f = 1/(3*float64(x)*float64(x)+10) + 10*float64(y) + 7

	fmt.Println(f)
}
