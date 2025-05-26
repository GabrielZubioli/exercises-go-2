package main

import (
	"fmt"
	"strings"
)

func main() {
	var numeros []int
	var input string

	for {
		var num int
		fmt.Printf("Digite um número: ")
		fmt.Scanln(&num)
		numeros = append(numeros, num)

		fmt.Print("Deseja inserir mais um número? (S/N): ")
		fmt.Scanln(&input)

		input = strings.ToLower(strings.TrimSpace(input))
		if input == "n" {
			break
		} else if input != "s" {
			fmt.Println("Entrada inválida, assumindo 'N' para finalizar.")
			break
		}
	}

	if len(numeros) == 0 {
		fmt.Println("Nenhum número foi digitado.")
		return
	}

	var soma int
	for _, valor := range numeros {
		soma += valor
	}

	media := float64(soma) / float64(len(numeros))
	fmt.Printf("A média dos números informados é: %.2f\n", media)
}
