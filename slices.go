package main

import "fmt"

func main(){             //0,1 ,2,3 ,4,5 ,6,7,8,9
	var numeros = [10]int {1,29,3,44,5,66,7,8,9,15}

	//slice
	fmt.Println(numeros[2:4])
	fmt.Println(numeros[2:])
	fmt.Println(numeros[:4])
	fmt.Println(numeros[:])

	var nome string = "Anderson Tiago"
	fmt.Println(nome[:])
	fmt.Println(nome[5:8])
	fmt.Println(nome[7:])
	fmt.Println(nome[:5])
	fmt.Println("==========")

	var nomes = [3]string{
		"Ana",
		"Jose",
		"Maria",
	}
	fmt.Println(nomes[0:2])
	var p1 []string = nomes[0:1]
	fmt.Println(p1)
	fmt.Println(nomes[2:])
	fmt.Println(nomes[1:2])
	fmt.Println(nomes[1:3])
	fmt.Println("Array original: ",nomes)

	//Modificando parte do array mesmo sendo outra variável, afeta ao array original.
	p1[0] = "Tiago"
	fmt.Println(p1)	
	fmt.Println("Array original modificado: ",nomes)

}