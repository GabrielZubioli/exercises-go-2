package main

import (
    "fmt"
    "sort"
)

func main() {
    var numeros [12]int

    for i := 0; i < len(numeros); i++ {
        fmt.Printf("Digite o %d° numero: ", i+1)
        fmt.Scanln(&numeros[i])
    }

    numerosSlice := numeros[:] 
    sort.Ints(numerosSlice)

    fmt.Println("Números ordenados:", numerosSlice)
}
