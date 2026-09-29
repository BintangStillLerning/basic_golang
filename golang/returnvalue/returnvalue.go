package main

import "fmt"

func namaorang()(NamaAwal string, NamaAkhir string, Umur int,){
	NamaAwal = "Bintang"
	NamaAkhir = "Akasyah"
	Umur = 17
    
	return
}

func main (){
   a,b,c := namaorang()  //nama variablenya bebas asal parameternya semua ada atau di hiraukan
   fmt.Println(a,b,c)
}