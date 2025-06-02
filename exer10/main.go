package main

import "fmt"

func main(){
	var num int
	fmt.Print("Digite um numero : ")
	fmt.Scanln(&num)
	
	if num % 2 == 0{
		fmt.Printf("o numero %d é par", num)
	}else{
		fmt.Printf("O numero %d é impar", num)
	}
 }