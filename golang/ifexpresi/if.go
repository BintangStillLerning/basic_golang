package main

import "fmt"

func main (){
	
	var namasaya = "bintang akasyah"

	if namasaya == "bintang akasyah" {
		fmt.Println("halo", namasaya)
	}else if namasaya == "bintang akasyah mpruy" { //else itu gk boleh di enter (harus satu baris setelah }
		fmt.Println("Oh ini temennya mpruy")
	}else {
		fmt.Println("jawaban salah, Siapa kamu?")
	}

	var panjanghuruf = len(namasaya)

	if panjanghuruf > 10 {
		fmt.Println("jangan panjang panjang woi")
	} else if panjanghuruf < 10 {
		fmt.Println(" ya udh bener coy")
	} else {
		fmt.Println("udh gk bener, ini apaan?")
	}
}
