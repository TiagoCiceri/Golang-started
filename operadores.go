package main

import "fmt"

func main(){
	var a int = 23
	var b int = 7

	fmt.Printf("Valor de a: %v  e b: %v \n", a , b)	
	fmt.Println("a > b =", a > b)
	fmt.Println("a < b =", a < b)
	fmt.Println("b >= a =", b >= a)

	//igualdade
	fmt.Println("3 == 3", 3 == 3)
	fmt.Println("3 == 4", 3 == 4)
	fmt.Println("a != b =", a != b)
	fmt.Println("7 != 6 =", 7 != 6)
	fmt.Println("7 != 7 =", 7 != 7)

	//logicos
	var x int = 1 	
	fmt.Println("x=",x)
	fmt.Println("x > 3 && x < 7 =", x > 3 && x < 7)
	fmt.Println("x < 3 || x > 7 =", x < 3 || x > 7)
	x = 4
	fmt.Println("x=",x)
	fmt.Println("x > 3 && x < 7 =", x > 3 && x < 7)
	fmt.Println("x < 3 || x > 7 =", x < 3 || x > 7)
	

}