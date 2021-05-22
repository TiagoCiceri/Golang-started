package main

import (
	"fmt"
)

type Geometrica interface {
	area() float64
}

type Quadrado struct {
	lado float64
}// area = lado ^ 2  (lado * lado)

type Circulo struct {
	raio float64
}// area = pi * (raio ^2)  (raio * raio)

func (str Quadrado) area() float64 {
	return str.lado * str.lado
}

func main(){
	var iG Geometrica
	iG = Quadrado{3}

	fmt.Printf("A area do quadrado é %v\n",  iG.area())

}