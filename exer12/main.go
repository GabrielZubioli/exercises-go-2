package main
import "fmt"

func main(){
	var num1, num2 int
	fmt.Print("Digite o primeiro numero : ")
	fmt.Scanln(&num1)
	fmt.Print("Digite o segundo numero : ")
	fmt.Scanln(&num2)

	if  (num1 + num2) % 2 == 0 {
		fmt.Printf("O resultado da divisao é zero")
	}else{
		fmt.Print("o resultado da divisao nao é zero")
	}
}