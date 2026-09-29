package main

import "fmt" 


func main (){

	for hitungan := 1; hitungan<=10; hitungan++ {
		fmt.Println("Perulangaan ke", hitungan)
	}

	slice := []string{"Bintang", "PM", "Ropi", "Paypay", "Mrpuy1"}
	for i := 0; i < len(slice); i++ {
	  fmt.Println("halo", slice[i])
	}
	// jadi slice := []tipedata{"data1", "data2"},
	//  for variablenya terus nilainya ; 
	// variable lebih besar atau lebih kecil dari panjang slice ;
	//  increment atau decrement variablenya
	// fmt.Println("halo", slicenya[variable yang tadi di forloop])
}