package main

import "fmt"

func somar(a float64, b float64) float64 {
	return a + b
}

func cumprimentar(nome string) func()string{
	return func() string { 
		return fmt.Sprintf("Olá, %s!!!", nome)
	}
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

	//Função que retorna uma função
	cumprimentarJoao := cumprimentar("João")
	cumprimentarTiago := cumprimentar("Tiago")

	fmt.Println(cumprimentarJoao)
	fmt.Println(cumprimentarTiago)
	fmt.Println(cumprimentarJoao)
	fmt.Println(cumprimentarTiago)
	fmt.Println(cumprimentarJoao)
	fmt.Println(cumprimentarTiago)
	fmt.Println(cumprimentarJoao)
	fmt.Println(cumprimentarTiago)

}