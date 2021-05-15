package main

import "fmt"

func main(){
	var soma int 
	
	
	for i := 1; i <= 1000; i = i + 1 {
		soma = (soma + i)
	} 
	fmt.Println("Loop For... Soma = ", soma)
	
	soma = 0

	for i := 1; i <= 1000; i++ {
		soma += i
	} 
	fmt.Println("Loop For simplificado... Soma = ", soma)


}