package main

import "fmt"

func mpruy(nama string)string {
	return nama + "Jangan Amping"
}

func main (){
	mike := mpruy
		fmt.Println(mike("Mpruy "))
		fmt.Println(mpruy("Mpruy amat"))
	
}