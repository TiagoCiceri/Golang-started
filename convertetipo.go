package main

import (
	   "fmt" 
	   "time"
	   )

func main(){
	var a int = 41
	var b float64 = 19.7
	data := time.Now()
	

	// abaixo a conversão
	var c int = int(b) 
	var d float64 = float64(a)
	var descricao string
	descricao = data.Format("MM-DD-YYY")

	fmt.Println("O valor de a é :", a)
	fmt.Println("O valor de b é :", b)
	fmt.Println("O valor de c é :", c)
	fmt.Println("O valor de d é :", d)
	fmt.Println("O valor de data é :", data)
	//verificar tipo
	fmt.Printf("Tipo: %T valor: %v\n",data,data)
	fmt.Printf("Tipo: %T valor: %v\n",d,d)

	fmt.Println("O valor de descrição é :", descricao)
}
