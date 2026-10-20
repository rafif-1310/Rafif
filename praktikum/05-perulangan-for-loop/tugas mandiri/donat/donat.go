package main

import (
	"fmt"
)

func main() {
	var n, m, x, y int

	fmt.Scan(&n, &m, &x, &y)

	jumlahDonat := 0

	for {

		if n >= x && m >= y {
			n -= x
			jumlahDonat++
		} else {

			break
		}
	}

	fmt.Println(jumlahDonat)
}
