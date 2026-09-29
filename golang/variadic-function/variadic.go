package main

import "fmt"
             //variable angka akan berubah menjadi slice
func hitungan (angka ...int ) int{
  total := 0 //variable untuk menyimpan data hitungan
  for _, nilai := range angka{ // for itu untuk tipe for range dengan cara for index, variable := range (ambil variable parameter)
  total += nilai //disini jdi total tampungan data hitungan akan di tambahkan oleh range nilai 
  // yang akan di isi di hitungan bawah
  }
return total //setelah itu kita return variable ang menyimpan data hitungan
}

func main (){
	total := hitungan(10,15,20,30,50) //arraynya jadi bisa di masukin kayak parameter dengan menggunakan vararg 
	// (Variable Argumen)
	fmt.Println(total)

	slice := [] int {10,12,13,14,15} //jika ada data slice maka bisa langsung di masukina ke argumen di atas
	total = hitungan(slice...) //dengan cara total yang menyimpan data tadi di sama dengakan function si argumen di atas
	//lalu tutup kurung (variable yang menampung slicenya)
	fmt.Println(total) //setelah itu di print dengan variable total
}

