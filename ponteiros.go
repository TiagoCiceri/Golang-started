package main

import "fmt"

func main(){
	var a int = 32;
	var p *int;//<- ponteiros

	a = 32
	fmt.Println(a)
	p = &a
	fmt.Println(p)
	fmt.Println(*p)
	*p = 41
	fmt.Println(a)
	fmt.Println(*p)
	*p = *p + 10
	fmt.Println(a)
	fmt.Println(*p)

}