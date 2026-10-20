package main

import (
	"fmt"
)

func main() {
	var poin int
	totalPoin := 0

	for {

		fmt.Scan(&poin)

		if poin == 0 {
			break
		}

		totalPoin += poin
	}

	fmt.Println(totalPoin)
}
