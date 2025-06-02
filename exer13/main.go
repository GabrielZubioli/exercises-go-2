package main

import "fmt"

type Pessoa struct{
	Nome string
	Idade int
}

func main(){
	var nome string
	var idade int

	fmt.Print("Qual o seu nome : ")
	fmt.Scanln(&nome)
	fmt.Print("Qual a sua idade :")
	fmt.Scanln(&idade)

	pessoa := Pessoa{
		Nome: nome,
		Idade: idade,
	}
	

	fmt.Println("Nome: ", pessoa.Nome)
	fmt.Println("Idade: ", pessoa.Idade)
}