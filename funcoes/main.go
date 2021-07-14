package main

import "fmt"

func main() {
	/*
	resultado := soma(10,20)
	fmt.Println(resultado)
	*/

	/*
	resultado := somaTudo(3,5,6,7)
	fmt.Println(resultado)
	*/

	resultado := func(x ...int) func() int {
		res := 0

		for _, v := range x{
			res += v
		}
		return func() int {
			return res * res
		}
	}
	fmt.Println(resultado(1,2,3,4)())

	carro := Carro{
		Nome: "Nivus",
	}
	carro.andar()
}

func soma(a int, b int) (result int) {
	result = a + b
	return
}

func somaTudo(x ...int) int {
	resultado := 0
	for _,v := range x {
		resultado += v
	}
	return resultado
}

// POO
type Carro struct {
	Nome string
}

func (c Carro) andar() {
	fmt.Println(c.Nome, "andou")
}