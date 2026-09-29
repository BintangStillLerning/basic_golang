package main

import (
	"errors"
	"fmt"
)


func pembagian (hasil int, nilai int)(int, error){
	if nilai == 0{
		return 0, errors.New("Pembagian 0 Gk bisa berakhir")
	} else{
		result := hasil / nilai
        return result, nil
	}
}

func main (){
	result, err := pembagian(100,0)
	if err == nil{
		fmt.Println("hasilnya", result)
	} else {
		fmt.Println("ini eror", err.Error())
	}
	// var hasilnilai error = errors.New("EROR WOI")
	// fmt.Println(hasilnilai)
}

