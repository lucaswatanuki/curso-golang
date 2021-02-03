package main

import (
	"fmt"
	"time"
)

func main() {

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	pessoa1 := Pessoa{"Joao"}
	pessoa2 := Pessoa{"Joao"}

	fmt.Println("Pessoas iguais ->", pessoa1 == pessoa2)
}
