package main

import "fmt"
import "time"

func main(){
	var soma int = 2

	for ; soma < 600 ; {
		soma = soma + soma
	}

	fmt.Println("Loop while... soma = ", soma)

	//reset	
	soma = 2

	for soma < 600 {
		soma += soma
	}

	fmt.Println("Loop while simplificado... soma = ", soma)

	//loop infinito
	for true {
		soma += soma
		fmt.Println(soma)
		time.Sleep(100 * time.Millisecond)
		if (soma > 10000000){
			break
		}
	}

}