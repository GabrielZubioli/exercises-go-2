package main

import (
	"fmt"
	"strings"
)

func main() {
	var moedaOrigem, moedaDestino string
	var quantidade float64

	for {
		fmt.Print("Qual a moeda de origem? (Dollar, Real, Euro): ")
		fmt.Scanln(&moedaOrigem)
		moedaOrigem = strings.ToLower(strings.TrimSpace(moedaOrigem))
		if moedaOrigem == "dollar" || moedaOrigem == "real" || moedaOrigem == "euro" {
			break
		} else {
			fmt.Println("Entrada inválida.")
		}
	}

	for {
		fmt.Print("Qual a moeda de destino? (Dollar, Real, Euro): ")
		fmt.Scanln(&moedaDestino)
		moedaDestino = strings.ToLower(strings.TrimSpace(moedaDestino))
		if moedaDestino == "dollar" || moedaDestino == "real" || moedaDestino == "euro" {
			break
		} else {
			fmt.Println("Entrada inválida.")
		}
	}

	fmt.Print("Qual a quantidade a ser convertida? ")
	fmt.Scanln(&quantidade)

	var dollar float64 = 5.67
	var euro float64 = 6.12
	var valorEmReais float64

	switch moedaOrigem {
	case "dollar":
		valorEmReais = quantidade * dollar
	case "euro":
		valorEmReais = quantidade * euro
	case "real":
		valorEmReais = quantidade
	}

	var resultado float64
	switch moedaDestino {
	case "dollar":
		resultado = valorEmReais / dollar
	case "euro":
		resultado = valorEmReais / euro
	case "real":
		resultado = valorEmReais
	}

	fmt.Printf("%.2f %s equivalem a %.2f %s\n", quantidade, strings.Title(moedaOrigem), resultado, strings.Title(moedaDestino))
}
