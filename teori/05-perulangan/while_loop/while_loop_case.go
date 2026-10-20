package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	loop := 1
	for loop <= n{
				fmt.Print(loop," ")
		loop++
	}
}
