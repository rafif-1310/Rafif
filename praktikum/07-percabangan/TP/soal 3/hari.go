package main

import "fmt"

func main(){
	var hari int

	
	fmt.Scan(&hari)
	
	switch hari {
	case 1:
		fmt.Print("senin")
	case 2:
		fmt.Print("selasa")
	case 3:
		fmt.Print("rabu")
	case 4:
		fmt.Print("kamis")
	case 5:
		fmt.Print("jumat")
	case 6:
		fmt.Print("sabtu(akhir pekan)")
	case 7:
		fmt.Print("minggu(akhir pekan)")
	default:
		fmt.Print("nomot tidak valid")
	}
		
	

}