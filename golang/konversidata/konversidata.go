package main

import "fmt"

func main () {
 
	var nilai8 uint8 = 129
	var nilai9 uint32 = 32323232
	var nilai10 uint8 = uint8(nilai9)
	var nilai11 uint16 = 32323
	var nilai12  uint8 = uint8(nilai11)


	fmt.Println( nilai8)
	fmt.Println( nilai9)
	fmt.Println( nilai10)
	fmt.Println( nilai11)
	fmt.Println( nilai12)

	var mpruy = "Anjay"
	var mpruy2 byte = mpruy[9]
	var mpruy3 string = string(mpruy2)
	fmt.Println(mpruy3)

}