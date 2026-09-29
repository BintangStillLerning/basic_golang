package main

import "fmt"

func main () {

	var namaorang [6]string

	namaorang[0] = "bintangakasyah"
	namaorang[1] = "adit"
	namaorang[2] = "dewa"
	namaorang[3] = "wowo"
	namaorang[4] = "mpruy"
	
	fmt.Println("Nama orang pertama adalah:", namaorang[0])
	fmt.Println("Nama orang kedua adalah:", namaorang[1])
	fmt.Println("Nama orang ketiga adalah:", namaorang[2])
	fmt.Println("Nama orang keempat adalah:", namaorang[3])
	fmt.Println("Nama orang kelima adalah:", namaorang[4])

	//atau juga bisa langsung seperti ini
	var namamobil = [10]string{
		"toyota",
		"honda",
		"nissan",
		"suzuki",
		"mazda",
	}
	fmt.Println(namamobil)

	//menghitung panjang array
	var panjangarray = len(namamobil)
	fmt.Println("Panjang array adalah:", panjangarray)

	//kalo array int
	var nilai1 = [3]int{

		(80),		
		(38),
		(40),
	}
	fmt.Println("nilai1 adalah:", nilai1)
}