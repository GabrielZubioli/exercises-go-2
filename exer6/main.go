package main

import (
	"fmt"
)

type Domino struct {
	A int
	B int
}

func contarGraus(dominos []Domino) map[int]int {
	graus := make(map[int]int)
	for _, d := range dominos {
		graus[d.A]++
		graus[d.B]++
	}
	return graus
}

func podeFormarCadeia(graus map[int]int) bool {
	for _, grau := range graus {
		if grau%2 != 0 {
			return false
		}
	}
	return true
}

func construirCadeia(atual int, dominos []Domino, usados []bool, caminho []Domino) ([]Domino, bool) {
	if len(caminho) == len(dominos) {
		if caminho[0].A == caminho[len(caminho)-1].B ||
			caminho[0].A == caminho[len(caminho)-1].A {
			return caminho, true
		}
		return nil, false
	}

	for i, d := range dominos {
		if usados[i] {
			continue
		}

		if d.A == atual || d.B == atual {
			usados[i] = true
			proximo := d.B
			if d.A == atual {
				proximo = d.B
				caminho = append(caminho, d)
			} else {
				proximo = d.A
				caminho = append(caminho, Domino{A: d.B, B: d.A})
			}

			if resultado, ok := construirCadeia(proximo, dominos, usados, caminho); ok {
				return resultado, true
			}

			usados[i] = false
			caminho = caminho[:len(caminho)-1]
		}
	}

	return nil, false
}

func main() {
	dominos := []Domino{
		{2, 1},
		{2, 3},
		{1, 3},
	}

	graus := contarGraus(dominos)

	if !podeFormarCadeia(graus) {
		fmt.Println("Não é possível formar uma cadeia de dominós cíclica.")
		return
	}

	for numero := range graus {
		usados := make([]bool, len(dominos))
		caminho, ok := construirCadeia(numero, dominos, usados, []Domino{})
		if ok {
			fmt.Println("Cadeia de dominós encontrada:")
			for _, d := range caminho {
				fmt.Printf("[%d|%d] ", d.A, d.B)
			}
			fmt.Println()
			return
		}
	}

	fmt.Println("Não foi possível formar uma cadeia de dominós cíclica.")
}
