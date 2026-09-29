package main

import "fmt"

    //func yang ini #buat penjelasan paling bawah
func filterkatakasar (kata string, filter func(string)string){
 katayangdifilter := filter(kata)
  fmt.Println("Hallo user", katayangdifilter)
}

func filteran(kata string)string {

	if kata == "Anjing"{
		return "Jangan Ngomong kasar yah"
	}else {
		return kata
	}

}

func main(){
    filterkatakasar("Aku suka main rp", filteran)// jadi dipanggil lagi di func main, ()di isi func awal yang buat filterkatakasar
	// caranya namafuncnya(di atas kan kata string, jadi kata yang mau di masukin itu buat di filter,
	// lalu masukan func filterannya) jadi filterkatakasar("kata yang mau di masukin, func yang mengandung argumen filteran)
}
