package main

import "fmt"

type Perfil struct {
	Altura  float64
	Ativo bool
	Profissao string
}

func main(){
	var perfis map[string]Perfil = make(map[string]Perfil)

	perfis["João"] = Perfil{
		1.74, true, "Médico",
	}
	
	perfis["Maria"] = Perfil{
		1.59, false, "Advogada",
	}
	fmt.Println(perfis)

}