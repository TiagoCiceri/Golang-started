package main

import "fmt"

type Pessoa struct {
	Nome string
	Sobrenome string
}

func (str Pessoa) nomeCompleto() string {
	return str.Nome +" "+ str.Sobrenome
}

func nfNomeCompleto(str Pessoa) string{
	return str.Nome+" "+str.Sobrenome
}

func main(){
	alguem := Pessoa{"João", "de Alencar"}

	fmt.Println(alguem)
	fmt.Println("Nome completo: ", alguem.nomeCompleto())
	fmt.Println("Nome completo: ", nfNomeCompleto(alguem))
}