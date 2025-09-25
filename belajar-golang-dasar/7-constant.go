package main

import "fmt"

func main(){
	const(
		firstname string = "Drajat"
		lastname  = "Wibowo"
	)
	//error
	//firstname ="Tidak Bisa di ubah"
	//lastname ="Tidak Bisa di ubah"

	fmt.Println(firstname)
	fmt.Println(lastname)
}