package main

import "fmt"

func main(){
	var a int = 35

	fmt.Println("a=",a)

	if (a > 50) {
		fmt.Println("a é maior que 50")
	}else if (a >= 20){
		fmt.Println("a esta entre 20 e 50")
	}else{
		fmt.Println("a é menor ou igual 50")
	}
}