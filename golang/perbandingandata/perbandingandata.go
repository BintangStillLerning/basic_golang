package main

import "fmt"

func main () {

	var nilaiA  = 25000
	var nilaiB  = 30000

	var HasilPerbandingan bool  = nilaiA > nilaiB //lebih besar dari
	var HasilPerbandingan2 bool = nilaiA < nilaiB //lebih kecil dari
	var HasilPerbandingan3 bool = nilaiA == nilaiB //sama dengan
	var HasilPerbandingan4 bool = nilaiA != nilaiB //tidak sama dengan
	var HasilPerbandingan5 bool = nilaiA >= nilaiB //lebih dari sama dengan
	var HasilPerbandingan6 bool = nilaiA <= nilaiB //kurang dari sama dengan

	fmt.Println("Hasil Perbandingan adalah:", HasilPerbandingan)
    fmt.Println("Hasil Perbandingan adalah:", HasilPerbandingan2)
	fmt.Println("Hasil Perbandingan adalah:", HasilPerbandingan3)
	fmt.Println("Hasil Perbandingan adalah:", HasilPerbandingan4)
	fmt.Println("Hasil Perbandingan adalah:", HasilPerbandingan5)
	fmt.Println("Hasil Perbandingan adalah:", HasilPerbandingan6)

	//atau juga bisa langsung seperti ini
	fmt.Println(nilaiA > nilaiB)
	fmt.Println(nilaiA < nilaiB)
	fmt.Println(nilaiA == nilaiB)
	fmt.Println(nilaiA != nilaiB)
	fmt.Println(nilaiA >= nilaiB)
	fmt.Println(nilaiA <= nilaiB)
}