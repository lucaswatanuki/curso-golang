package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	//int para string
	fmt.Println("Valor em string ->" + strconv.Itoa(notaFinal))

	//string para int
	num, _ := strconv.Atoi("123456")
	fmt.Println("Tipo da variável ->", reflect.TypeOf(num))
	fmt.Println("Valor ->", num)

}
