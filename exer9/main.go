package main

import "fmt"

func main(){
	var idade int
	var nome string

	fmt.Print("Qual o seu nome :")
	fmt.Scanln(&nome)

	fmt.Print("Qual a sua idade :")
	fmt.Scanln(&idade)

	fmt.Printf("Seu nome é %s e você tem %d anos", nome, idade)
}