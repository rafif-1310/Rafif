package main

import "fmt"

func main() {
	var p, q int
	
	fmt.Scan(&p, &q)
	hasil1 := (p%2 == 0) || (q%2 == 0)
	hasil2 := (p%2 != 0) && (q%2 != 0)
	hasil3 := !(p == q)

	fmt.Printf("%t %t %t\n", hasil1, hasil2, hasil3)


}