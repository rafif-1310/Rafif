package main

import "fmt"

func main() {
	//Cara 1
	var name string
    name = "Rafif Ramadhan"

	fmt.Println("Name :", name)
     //Cara 2
	var lastName string = "Ramadhan"
	fmt.Println("Last Name :", lastName)
    //Cara 3
	var middleName = "Rafif"
	fmt.Println("middleName :", middleName)
   // Cara 4
	var (
		fullname = "Rafif Ramadhan"
		firstname = "Rafif"
	)
	fmt.Println(fullname)
	fmt.Println(firstname)

	
}
