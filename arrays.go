package main

import "fmt"

func main(){	
	var numeros [7] int
	
	for i := 0; i < 7; i++ {
		numeros[i] = i
	}
	
	fmt.Println(numeros)	
}