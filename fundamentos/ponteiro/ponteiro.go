package main

import "fmt"

func main() {
	i := 1

	//Referenciando memória
	var p *int = nil

	p = &i //pegando endereço de memória da variável 'i'

	//Apontando para onde 'p' aponta
	var y *int = p

	i++

	fmt.Println("Endereço de memoria -> ", p)
	fmt.Println("Ponteiro y -> ", *y)
	fmt.Println("Ponteiro p -> ", *p)
	fmt.Println("Valor de i -> ", i)

}
