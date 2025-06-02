package main

import (
	"fmt"
)

type Pessoa struct{
	Nome string
	Idade int
	Email string
}

func main(){
	var  nome, email, criarMais string
	var idade int
	var pessoas []Pessoa

	for true{

		fmt.Print("Digite o nome: ")
		fmt.Scanln(&nome)

		fmt.Print("Qual a idade: ")
		fmt.Scanln(&idade)

		fmt.Print("Digite o email: ")
		fmt.Scanln(&email)

		pessoa := Pessoa{
			Nome : nome,
			Idade : idade,
			Email : email,
		}

		pessoas = append(pessoas, pessoa)
		fmt.Print("Desenha criar mais uma pessoa ? ")
		fmt.Scanln(&criarMais)
		if criarMais == "Não" {
			break
		}
	}
	for i, pessoa := range(pessoas){
		fmt.Printf("%d - %s, %d, %s\n", i+1, pessoa.Nome, pessoa.Idade, pessoa.Email)
	}
}