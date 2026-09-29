package main 

import "fmt"

func Hallo (nama string) string{
	if nama == ""{
		return "Hello Brow"
		}else{
			return " Halloo " + nama
		}
		//return itu biasanya harus di taro bawah, kareba setelah return dia tidak akan mengeksekusi 
		// jawaban di bwahnya lagi, jadi kalo gk ditaro bawah palingan di buat tutup kurung awal 
		// buat membuat return baru kayak if else atau for loops
		
	}



func main (){
   hasil := Hallo("")
   fmt.Println(hasil)
}