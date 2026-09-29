package main

import "fmt"

type Alamat struct{
	alamat, jalan, kota string
}
func main(){
	alamat1 := Alamat{"j;.Pisang","anjay",  "Jl.Mpruy"}//pointer berfungsi sebagai operator yang 
	//mengubah value di variable yang berbeda tapi sama 
	alamat2 := &alamat1 //jadi kalo ente make pointer(&) ke alamat1
	//otomatis alamat2 adalah alamat1
    alamat3 := &alamat1
	// *alamat4 = &Alamat{"jakarta", "bandung", "Ngawi"}

//tanda & ini adalah pointer
    alamat2.alamat = "Bandung"//alamat2 alias alamat satu hanya mengubah parameter pertama
	//yang sesuai dengan urutan parameternya
	*alamat2 = Alamat{"jakarta", "bandung", "Ngawi"}//kalo ini alamat2 dipointer ke alamat
	//dengan begitu alamat2 akan membuat data baru dengan data alamat1 

	var alamat4 *Alamat= new(Alamat) //dengan function new maka data baru akan terbuat namun kosong
	// alamat4.alamat = "Jakarta" | kalo mau nambahin data baru ke data kosong 
    //ketika di jalankan maka operator * mengubah semua data yang di *alamat2 ke data Alamat
	//jadi hasilnya sama semua karena di pindahkan ke alamat3 
	fmt.Println(alamat1)
	fmt.Println(alamat2)
	fmt.Println(alamat3)
	fmt.Println(alamat4)
}