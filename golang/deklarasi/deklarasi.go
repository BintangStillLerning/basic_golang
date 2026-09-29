package main

import "fmt"

func main() {

    type datanyaApa int
	var umurku datanyaApa = 23
	var umurtemen datanyaApa = 20
	fmt.Println("Umur adek saya adalah:", umurtemen)
	fmt.Println("Umur saya adalah:", umurku)

    type NoKTP string //Deklarasi tipe data baru dengan alias
	var Noktpaku NoKTP = "1092102120192"
	var Noktptemen NoKTP = "1930192390230"
	fmt.Println("Nomor KTP temen saya adalah:", Noktptemen)
	fmt.Println("Nomor KTP saya adalah:", Noktpaku)

	type statusNikah bool
	var StatusNikahTemen statusNikah = true
	var statusNikahSaudara statusNikah = false
	fmt.Println("Status Nikah saudara saya adalah:", statusNikahSaudara)
	fmt.Println("Status Nikah temen saya adalah:", StatusNikahTemen)
}