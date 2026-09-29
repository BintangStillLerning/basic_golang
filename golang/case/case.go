package main

import "fmt"

func main () {
    
	fmt.Println("Kamu suka makan apa??")
	makanan := "mie ayam"
	fmt.Scanln(&makanan)

	switch makanan {
	case "sate": fmt.Println("ini sih enak ya apa lagi kalo sate kambing")
	case "nasi goreng": fmt.Println("wah ini mah enak banget")
	case "mie ayam": fmt.Println(" Ter the Best ini mie aya gatot ")
	default: fmt.Println("baru denger, makanan apaan tuh")
	}
}