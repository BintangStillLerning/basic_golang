package main

import (
	"fmt"
)

type Datapengguna struct{
	Nama string
	Alamat string
	No_telepon int
	Umur int
}

func main(){
	var Udin Datapengguna
	Udin.Nama = "Bintang Akasyah"
	Udin.Alamat = "Jl.Kebo Mutar Keliling"
	Udin.Umur = 18

	fmt.Println(Udin.Nama)
	fmt.Println(Udin.Alamat)
	fmt.Println(Udin.Umur)

	Joko := Datapengguna{
	Nama : "BintangDkp",
	Alamat : "Jl.Buah Komplek Demikian",
	Umur : 35,
}

 fmt.Println(Joko)
}

