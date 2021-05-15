package main

import "fmt"

func main(){
	var nomes = []string {
		"Ana",
		"Jose",
		"Maria",
	}

	fmt.Println(nomes)
	//
	fmt.Printf("Capacidade cap=%d\n", cap(nomes))
	fmt.Printf("Tipo: %T\n", nomes)
	fmt.Println("Tamanho len=", len(nomes))
	fmt.Printf("cap=%d --- len=%d\n", cap(nomes), len(nomes))
	fmt.Printf("cap=%d --- len=%d\n", cap(nomes[1:]), len(nomes[1:]))

	fmt.Println("----- Exercicio-----")
	//Exercicio
	var animes [3]string;
	animes =[3]string{"Dragon Ball", "Sailor Moon", "Yuyu Hakusho"}
	var parte = animes[0:2]
	fmt.Println(animes)
	fmt.Println(parte)
	fmt.Printf("cap=%d --- len=%d\n", cap(parte), len(parte))
}