package main

import "fmt"

func main () {

	const namavariable = "Tipe data bebas bisa string DLL"
	fmt.Println(namavariable)

	//contoh erornya

	//namavariable = "Coba diubah"
	//fmt.Println(namavariable)

	// gk bisa di ubah karena sudah di deklarasi sebagai constant

	const sprint = true
	fmt.Println("Apakah ini sebuah constant boolean?", sprint)
}