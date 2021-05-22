package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	Nome string
	Sobrenome string
}

//Metodo com receptor
func (str Pessoa) nomeCompleto() string {
	return str.Nome +" "+ str.Sobrenome
}

//Metodo sem receptor
func nfNomeCompleto(str Pessoa) string{
	return str.Nome+" "+str.Sobrenome
}

//Metodo com ponteiro receptor --- sem usar ponteiro a alteração é feita apenas na cópia
func (str *Pessoa) capitalizarNome (){
	str.Nome = strings.ToUpper(str.Nome)
}

//Metodo sem receptor com ponteiro NÃO FUNCIONA
func fnCapitalizarNome(str *Pessoa){ 
	str.Nome = strings.ToUpper(str.Nome)
}

func main(){
	alguem := Pessoa{"João", "de Alencar"}

	fmt.Println(alguem)
	fmt.Println("Nome completo: ", alguem.nomeCompleto())// com receptor

	alguem.capitalizarNome()

	fmt.Println("Nome completo: ", nfNomeCompleto(alguem))// sem receptor

}