package main

import (
	"fmt"
	
)

type blacklist func(string)bool // jadi function tanpa nama (anonymous) tinggal dikasi tipe datanya aja

func RegisterUser(user string, BlackList blacklist){ //func yang di buat untuk logika block user atau admin
 if BlackList (user){
	fmt.Println("You Has Been Banned Permanently", user)
 }else {
	fmt.Println("Welcome Back", user)
 }
}

// func blacklistuser (user string)bool{
// 	return user == "Mpruy"
// }


func main (){    //di isi buat variable yang menamping user yang di blacklist dan user yang ada
 blacklisted := func(user string)	bool{ //variable data user yang di blacklist
	return user == "Mike"
	} //ambil func logika block user lalu ("NamaUser", Variable data user yang di blacklist biar di cek func awal (blacklisted) )
	RegisterUser("admin", blacklisted) 
	RegisterUser("Mike", blacklisted)
}

