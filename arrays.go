package main

import "fmt"

func main(){		
	var numeros [17] int
	
	for i := 0; i < len(numeros); i++ {
		numeros[i] = i
	}	
	fmt.Println(numeros)	

	var novoarray = [5] int {10, 20, 30, 4, 5}	
	fmt.Println(novoarray)

	var soma int =0

	for i := 0; i < len(novoarray); i++ {
		soma += novoarray[i]
	}
	fmt.Println("Soma do array: ", soma)

}