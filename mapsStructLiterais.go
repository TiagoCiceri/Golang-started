package main

import "fmt"

type Perfil struct {
	Altura  float64
	Ativo bool
	Profissao string
}

func main(){
	var perfis map[string]Perfil = map[string]Perfil{
		"João": {
			1.74, true, "Médico",	
		},
		"Maria":{
			1.59, false, "Advogada",
		},
		"Fulano":{
			1.59, false, "Autonomo",
		},
	}

	
	fmt.Println(perfis)

}