package main 

import "fmt"


func Ups(i int) interface{}{
	if i == 1{
		return 1
	} else if 1 == 2{
		return true
	}else {
		return "ups"
	}

}

func main (){
	var hasil interface{} = Ups(23123)
		fmt.Println(hasil)
	
}
