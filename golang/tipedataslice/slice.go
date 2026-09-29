package main

import "fmt"

func main () {

	var bulan = [30]string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	var slicebulan = bulan[:]
	fmt.Println("Bulan pada slice adalah:", slicebulan)
	fmt.Println("panjang bulan adalah", len(slicebulan))
	fmt.Println("Kapasitas bulan adalah", cap(slicebulan))

	//menambah elemen atau atau data baru pada slice
    bulanbaru := append(slicebulan, "jember")
	slicebulan[0] = "janujanu"
	fmt.Println(bulanbaru)
	fmt.Println(slicebulan)

	slicebaru := make ([]string, 2, 5)
	fmt.Println(slicebaru)
	fmt.Println(len(slicebaru))
	fmt.Println(cap(slicebaru))
	
	
}