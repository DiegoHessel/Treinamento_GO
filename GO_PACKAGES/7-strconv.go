package main

import (
	"fmt"
	"strconv"
)

// convertendo string para inteiro
func convercaoDeStringParaNumero() {
	numeroStr := "123"

	// strconv.Atoi = parseint. Então ele vai converter uma string para inteiro.
	// Caso a string não seja um número válido, ele retornará um erro que deve ser tratado.
	numero, err := strconv.Atoi(numeroStr)
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
	fmt.Println("Número:", numero)
}

//converção de inteiro para string
func convercaoDeNumeroParaString(){
	numero2 := 4566
	// O método strconv.Itoa converte um número inteiro em uma string.
	// Ele é útil para transformar valores numéricos em texto para exibição ou manipulação.
	numeroStr2 := strconv.Itoa(numero2)
	fmt.Println("Número:",numeroStr2 )
}
func main() {
	convercaoDeNumeroParaString()
	convercaoDeStringParaNumero()
}
