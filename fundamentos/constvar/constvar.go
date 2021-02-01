package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415

	var raio = 3.2

	area := PI * math.Pow(raio, 2)

	fmt.Println("Área: ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 5
		f = "teste"
	)

	fmt.Println(a, b, c, d, f)

	var e, g bool = true, false
	fmt.Println(e, g)
}
