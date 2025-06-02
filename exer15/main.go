package main

import "fmt"

func main(){
	pessoas := []string {"gabriel", "amanda", "pedro","joao"}

	for i, pessoa := range(pessoas){
		fmt.Printf("%d - %s\n", i+1, pessoa)
	}
}