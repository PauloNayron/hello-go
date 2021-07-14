package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome 	string `json:"nome"`
	Email 	string `json:"email"`
}

func (c Cliente) toString() {
	fmt.Printf("{Nome: %s, Email: %s}\n", c.Nome, c.Email)
}

type ClienteBrasileiro struct {
	Cliente
	CPF 	int `json:"cpf"`
}

func main() {
	cliente := Cliente{"Paulo Nayron","paulo.nayron@jaya.tech"}
	cliente.toString()
	fmt.Println(cliente)
	fmt.Println("-------------------------------")
	clienteBrasileiro := ClienteBrasileiro{Cliente {"João", "joao@joao.com"}, 123452343}
	fmt.Printf("Nome: %s. Email: %s. CPF: %d\n", clienteBrasileiro.Nome, clienteBrasileiro.Email, clienteBrasileiro.CPF)
	clienteBrasileiro.toString()

	fmt.Println("-------------- exemplo de serialização em JSON -----------------")
	//	exemplo de serialização em JSON
	clienteJson, err := json.Marshal(clienteBrasileiro)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(clienteJson))

	fmt.Println("-------------- exemplo de deserialização -----------------")
	// exemplo de deserialização
	jsonBody := `{"nome":"João","email":"joao@joao.com","cpf":123452343}`
	clienteDeserializado := ClienteBrasileiro{}
	json.Unmarshal([]byte(jsonBody), &clienteDeserializado)
	fmt.Println(clienteDeserializado)
}
