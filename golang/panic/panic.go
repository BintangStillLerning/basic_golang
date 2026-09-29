package main

import "fmt"

func EndApp (){
	massage :=  recover()
		if massage != nil{ //recover berfungsi sebagai counter dari panic, 
		// sehingga aplikasi atau program tetap berjalan tanpa berhenti karena panic
			fmt.Println("Telah Terjadi Eror Harap Refresh Kembali: ","CRASH 9323", massage)
		}
	
	fmt.Println("Aplikasi Telah Selesai Berjalan")
}

func RunApp (error bool){
	defer EndApp()
	if error {
	   panic("ERORRR") //panic berfungsi ketika ada kesalahan atau eror dia akan otomatis menghentikan proses aplikasi 
	   //tanpa menjalankan function dibawahnya lagi(kecuali defer dan recover)
	}
	// Jangan Menambahkan Recover DIbawah atau setelah panic karena program recover tidak akan berjalan
	fmt.Println("Aplikasi Berjalan")
}

func main (){
	RunApp(false)
    fmt.Println("Program akan terus berjalan...") // program yang baris ini akan berjalan jika ada recover
}
