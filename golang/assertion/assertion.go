package main

import "fmt"

func random() interface{}{
		return 100
}

func main (){
	result := random()
	// resultString := result.(int)
	// fmt.Println(resultString)

	// resultInt := result.(int) //cara ini lebih ringkas tapi maksa dan susah di ingat
	// fmt.Println(resultInt)

	switch value := result.(type){
	case string:
		fmt.Println(" Ini berubah", value, "Menjadi") //lebih mudah dan paham cara ini
    //dengan switch case
	case int:
		fmt.Println("ini adalah", value, "Menjadi")
	
    default:
		fmt.Println("gk tau")
	}
	
}