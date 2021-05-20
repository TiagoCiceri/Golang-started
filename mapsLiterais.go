package main

import "fmt"

func main(){
	var notas map[string]int	

	notas = map[string]int{
		"Tiago" : 5,
		"Priscila": 10,
		"Tedy": 9,
		"Nina": 8,
		"Flor": 7,
		"Tigrão": 6,
	}

	fmt.Println(notas)
}