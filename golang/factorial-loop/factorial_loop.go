package main 

import "fmt"

func factorial (nilai int)int{
	hasil := 1
	for i := nilai; i > 0; i--{
		hasil *= i
	}
	return hasil
}

func main (){
	loop := factorial(5)
	fmt.Println(loop)
	recursive := factorialrecursive(10)
	fmt.Println(recursive)
	
}

func factorialrecursive(value int)int{
	if value == 1{
	return 1
	}else {
	return value * factorialrecursive(value-1) //cara ini lebih efektif dan singkat dibanding di atas
}
}