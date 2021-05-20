package main

import "fmt"

func somar(a float64, b float64) float64 {
	return a + b
}

func main(){
	//var funcaoSomar func(a float64, b float64) float64 = func(a float64, b float64) float64 {
	//	return a + b
	//}
	//abaixo forma reduzida
	funcaoSomar := func(a, b float64) float64 {
		return a + b
	}

	fmt.Println("Resultado soma:",somar(5, 6))

	fmt.Println("Resultado soma func variável:",funcaoSomar(5, 6))

	multiplicar := func(a,b float64) float64 {
		return a * b
	}

	fmt.Println("Resultado multiplicar:",multiplicar(5, 6))
}