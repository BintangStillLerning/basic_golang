package main

import "fmt"

func main (){

	for i := 0; i < 10; i++{
		if i % 3 == 0 {
			continue
		} 
// jadi kalo i dibagi 3 itu 0 maka akan di skip oleh statement continue
			fmt.Println("ini Perulangan ke ", i )
		
	}
}