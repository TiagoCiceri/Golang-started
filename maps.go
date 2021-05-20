package main

import "fmt"

func main(){
	var (
		valorx int
		existex bool
	)
	var notas map[string]int 

	notas = make(map[string]int)

	notas["Ana"] = 9
	notas["Maria"] = 10

	fmt.Println(notas)
	fmt.Println(notas["Maria"])

	valor, existe := notas["João"]

	if (existe){
		fmt.Println(valor)		
	}	
	
	valorx, existex = notas["Ana"]
	if (existex){
		fmt.Println(valorx)		
	}
}