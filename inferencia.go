package main

import "fmt"

func main(){
	var i int
	var j = i
	fmt.Printf("Tipo: %T ... Valor: %v\n", j, j)

	var x = 0.65 + 0.77
	fmt.Printf("Tipo: %T ... Valor: %v\n", x, x)	
}
