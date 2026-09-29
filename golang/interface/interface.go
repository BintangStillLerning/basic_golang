// package main

// import "fmt"

// type Kontrak interface{
// 	GetName() string
// }

// func DealKontrak(pelatih Kontrak){
// 	fmt.Println("Halloo",pelatih.GetName())
// }

// type Pelatih struct{
// 	Nama string
// 	Umur int
// }

// func (p Pelatih) GetName() string{
// 	return p.Nama
// }

// func main (){
//      mpruyjustin := Pelatih{
// 		Nama: "Coach Justin Mpruy",
// 		Umur: 64,
// 	 }
// 	var kontrakbaru Kontrak = mpruyjustin

// 	DealKontrak(kontrakbaru)
// }

package main

import "fmt"

type Kontrak interface{
	GetName() string
}

func DealKontrak (p Kontrak){
	fmt.Println("hallo",p.GetName())
}

type Pelatih struct{
	Nama string
}

func (k Pelatih) GetName()string{
	return k.Nama
}

func main (){
	Udin := Pelatih{
	Nama : "Udin Petot",
	}
	var kontrakbaru Kontrak = Udin

	DealKontrak(kontrakbaru)
}
    