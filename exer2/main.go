package main

import "fmt"

func main(){
	var nota1, nota2, nota3, nota4 int
	var media float64

	fmt.Print("Digite o valor da nota do primeiro bimestre\n")
	fmt.Scanln(&nota1)
	fmt.Print("Digite o valor da nota do segundo bimestre\n")
	fmt.Scanln(&nota2)
	fmt.Print("Digite o valor da nota do terceiro bimestre\n")
	fmt.Scanln(&nota3)
	fmt.Print("Digite o valor da nota do quarto bimestre\n")
	fmt.Scanln(&nota4)

	media = float64(nota1+nota2+nota3+nota4) / 4.0

	if (media >= 5 ){
		fmt.Printf("Aprovado !!\nValor da Média: %.2f", media)
	}else{
		fmt.Printf("Reprovado !!\nValor da Média: %.2f", media)
	}

}