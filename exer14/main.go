package main

import "fmt"

func criarCarro() Carro{
	var nome, placa string
	fmt.Print("Nome do carro: ")
	fmt.Scanln(&nome)
	fmt.Print("Placa do carro: ")
	fmt.Scanln(&placa)
	carro := Carro{
		Nome : nome,
		Placa : placa,
	}
	return carro
}

func mostrarGaragem(garagem []Carro){
	if len(garagem) == 0 {
		fmt.Print("Garagem vazia")
		return
	}

	for i, vaga := range garagem{
		fmt.Print("%d - %s (%s)\n",i+1, vaga.Nome, vaga.Placa)
	}
}

type Carro struct{
	Nome string
	Placa string
}

func main(){
var garagem []Carro
	var opcao int

	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1 - Criar carro")
		fmt.Println("2 - Mostrar garagem")
		fmt.Println("3 - Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			carro := criarCarro()
			garagem = append(garagem, carro)
			fmt.Println("Carro adicionado na garagem!")
		case 2:
			mostrarGaragem(garagem)
		case 3:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
	
		
}