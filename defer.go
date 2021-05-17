package main

import "fmt"

func main(){
	defer fmt.Println("Processo finalizado!!!")
	defer fmt.Println("Aguade...processo 1")
	defer fmt.Println("Aguade...processo 2")
	defer fmt.Println("Aguade...processo 3")

	fmt.Println("Aguade...")
}