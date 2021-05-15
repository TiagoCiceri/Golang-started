package main

import "fmt"

func calcular(a int) (int ,int){
	var quadrado int = a * a
	var cubo int = a * a * a

	return quadrado, cubo
}

func calcular2(a int) (quadrado int , cubo int){
	quadrado = (a * a)
	cubo = (a * a * a)	

	return
}

func main(){
	fmt.Println(calcular(2))
	fmt.Println(calcular2(2))
}