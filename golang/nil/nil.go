package main

import "fmt"

func NewMap (Nama string)  map [string]string{
	if Nama == ""{
		return nil
	}else {
		return map[string] string{
			"name": Nama,
		}
	}
}


func main (){
  var mpruy  map[string]string = NewMap("anjay")
  fmt.Println(mpruy)
}