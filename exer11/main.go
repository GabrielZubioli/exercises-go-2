package main 
import "fmt"

func main(){
	var num int

	fmt.Print("Digite um numero de 0 a 10: ")
	fmt.Scanln(&num)

	if num < 0 || num > 10{
		fmt.Print("Não é um numero de 1 a 10")
	}
}