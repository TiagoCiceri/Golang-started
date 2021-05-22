package main

import "fmt"


func acumulador() func(int) int {
	var soma int = 0
	
	return func(a int) int {
		soma += a
		return soma
	}
}

func main(){
	var add1 = acumulador()
	
	fmt.Println("A: ",add1(10))	
	fmt.Println("A: ",add1(12))	
	fmt.Println("A: ",add1(5))
	fmt.Println("--------------------------------")
	var add2 = acumulador()
	fmt.Println("B: ",add2(5))
	fmt.Println("B: ",add2(10))
	fmt.Println("--------------------------------")
	fmt.Println("A: ",add1(0))
	fmt.Println("B: ",add2(0))
}