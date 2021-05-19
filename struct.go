package main

import "fmt"

func main(){
	type Posicao struct{
		x int
		y int
	}

	pos := Posicao{40,67}
	fmt.Println(pos)
	pos.y = 51
	fmt.Println(pos)
	fmt.Println(pos.x)

	posic := Posicao{y: 77, x: 180}
	fmt.Println(posic)
}