package main

import "fmt"

func main(){
	var nomes []string;

	nomes = append(nomes, "Tiago")
	fmt.Println(nomes)

	nomes = append(nomes, "Priscila")
	fmt.Println(nomes)

	nomes = append(nomes, "Tedy","Nina","Flor","Tigrão")
	fmt.Println(nomes)

	fmt.Printf("len= %d ... cap= %d\n", len(nomes), cap(nomes))
}