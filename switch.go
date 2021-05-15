package main

import "fmt"

func main(){
	var nome string = "Tiago"

	switch nome {
	case "Priscila":
		fmt.Println("É a Priscila")
	case "Tiago":
		fmt.Println("É o Tiago")
	default:
		fmt.Println("Não conheço.")
	}
}