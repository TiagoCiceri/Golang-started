package main

import "fmt"

func main(){
	numeros := []int{10,20,30}

	fmt.Println("---------Loop For------------")
	for i := 0; i < len(numeros); i++{
		fmt.Println(numeros[i])
	}

	fmt.Println("---------Loop Range------------")		
	for _, v := range numeros {
		fmt.Println(v)
	}
	fmt.Println("---------ou---------")		
	for indice := range numeros {
		fmt.Println(numeros[indice])
	}
	
}