package main

import (
	"fmt"
	"strings"
)

func main(){
	tabuleiro := [][]string{
		[]string{"X", "-", "-"},
		[]string{"-", "X", "-"},
		[]string{"0", "0", "X"},
	}

	fmt.Println(tabuleiro)
	fmt.Println("--------------------------------")
	//
	for i := 0; i < len(tabuleiro); i++{
		fmt.Println(tabuleiro[i])
	}
	
	fmt.Println("--------------------------------")
	for i := 0; i < len(tabuleiro); i++{
		fmt.Printf("%s\n", strings.Join(tabuleiro[i], " "))
	}

	fmt.Println("--------------------------------")
	for i := 0; i < len(tabuleiro); i++{
		for j := 0; j < len(tabuleiro[i]); j++{
			valor := tabuleiro[i][j]
			fmt.Println("Linha:",i," Coluna:",j," Valor:",valor)	
		}
	}

	

}