package main

import "fmt"

func main(){
	var valor1, valor2, valor3 int
	fmt.Print("Digite o primeiro valor\n")
	fmt.Scanln(&valor1)
	fmt.Print("Digite o segundo valor\n")
	fmt.Scanln(&valor2)
	fmt.Print("Digite o terceiro valor\n")
	fmt.Scanln(&valor3)

	soma := valor1 + valor2 + valor3
	quadradoSoma := soma * soma
	
	fmt.Printf("O quadrado da soma dos três valores é: %d\n", quadradoSoma)

}