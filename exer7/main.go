package main

import "fmt"


func dobrar(x *int){
	*x = *x * 2
}
func dobrarSemP(x int) int{
	return x * 2
}
func main(){
	var num int
	fmt.Print("Digite um numero: ")
	fmt.Scanln(&num)
	// ele nao altera o valor da variacel
	num2 := dobrarSemP(num)
	dobrar(&num)
	fmt.Print("O numero dobrado com o uso do ponteiro é : ", num)
	fmt.Print("\nO numero dobrado sem o uso do ponteiro é : ", num2)

}