package main

import "fmt"

func main(){
	type Posicao struct {
		x int
		y int
	}

	var pos Posicao = Posicao{10 ,20}
	var ponteiro *Posicao = &pos

	fmt.Println(pos)
	fmt.Println(ponteiro)
	fmt.Println(ponteiro.x)
	ponteiro.x = 66
	fmt.Println(ponteiro.x)
	fmt.Println(ponteiro)
	fmt.Println(pos)

}