package main

import "fmt"

type Carro struct {
	Name string
}

func (c Carro) andou() {
	c.Name = "Cruze"
	fmt.Println(c.Name, "andou")
}

func (c *Carro) andouComPonteiro() {
	c.Name = "Cruze"
	fmt.Println(c.Name, "andou")
}

func main() {
	/*
	a := 10
	println("valor:", a)
	println("endereço:", &a)
	println("------------------------")
	var ponteiro *int = &a
	println("endereço:",ponteiro)
	println("valor:", *ponteiro)
	println("------------------------")
	*ponteiro = 50
	println("endereço:",ponteiro)
	println("valor:", *ponteiro)
	println("------------------------")
	println("valor:", a)
	println("endereço:", &a)
	*/
	/*
	variavel := 10
	abc(&variavel)
	fmt.Println(variavel)
	*/

	carro := Carro{
		Name: "Nivus",
	}
	carro.andou()
	fmt.Println(carro.Name)
	fmt.Println("----------------------")
	carro.andouComPonteiro()
	fmt.Println(carro.Name)
}

func abc(a *int) { *a = 200 }
