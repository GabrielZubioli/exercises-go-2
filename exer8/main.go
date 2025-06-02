package main

import "fmt"

func criarCarro(modelo, marca, nome string) map[string]string{
	carro := map[string]string{
		"modelo" : modelo,
		"marca" : marca,
		"nome" : nome,
	}
	return carro
}

func main(){
	var modelo, marca, nome string
	fmt.Print("Qual o modelo do carro : ")
	fmt.Scanln(&modelo)
	fmt.Print("Qual a marca do carro : ")
	fmt.Scanln(&marca)
	fmt.Print("Qual o nome do carro : ")
	fmt.Scanln(&nome)

	carro := criarCarro(modelo, marca, nome)
	
	fmt.Print("Esse é seu carro : \n")
	for chave, valor := range carro{
		fmt.Println(chave, " : ", valor)
	}
}