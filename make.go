package main

import "fmt"

func main(){
	numeros := make([]int, 5)
	fmt.Println(numeros)

	novonum := make([]int, 0, 5)
	fmt.Println(novonum)
	fmt.Println(novonum[0:2])
}