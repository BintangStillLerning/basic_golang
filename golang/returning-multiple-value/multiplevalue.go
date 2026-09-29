package main 

import "fmt"


func sibro()(string, string, int, string){
	return "bintang","akasyah", 15, "Tahun"
}

func main (){
	NamaAwal, NamaAkhir, Umur, Tahun := sibro()
		fmt.Println(NamaAwal, NamaAkhir, Umur, Tahun) 
	//si variable nama akhir atau apapun itu bisa di hiraukan dengan pake underscore (_)
	// dengan cara NamaAwal, _ := sibro()
	//dam di fmt.Println(NamaAwal) Hanya perlu variable yang ada di list karena si NamaAkhir Di hiraukan
}