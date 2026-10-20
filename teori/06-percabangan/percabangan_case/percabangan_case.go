package main

import "fmt"

func main(){
	lamp_status := "kunings"

	if lamp_status == "merah"{
		fmt.Println(" berhenti!")
	}else if lamp_status == "kuning"{
		fmt.Println("Hati-hati!")
	}else if lamp_status == "hijau"{
		fmt.Println("jalan!")
	}else{
		fmt.Println("Perintah tidak terdaftar disini")
	}
}