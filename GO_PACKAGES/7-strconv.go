package main

import (
	"fmt"
	"strconv"
)

// convertendo string para inteiro
func conversaoDeStringParaNumero() {
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

// converção de inteiro para string
func conversaoDeNumeroParaString() {
	numero2 := 4566
	// O método strconv.Itoa converte um número inteiro em uma string.
	// Ele é útil para transformar valores numéricos em texto para exibição ou manipulação.
	numeroStr2 := strconv.Itoa(numero2)
	fmt.Println("Número:", numeroStr2)
}

// O método strconv.ParseFloat converte uma string que representa um número decimal em um valor float64.
// Ele é útil para interpretar números decimais armazenados como texto. Caso a string não seja válida, retorna um erro.
func conversaoDeStringParaFoat() {
	floatStr := "12.9"

	valorFloat, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
	fmt.Println("Número:", valorFloat)
}

func main() {
	conversaoDeNumeroParaString()
	conversaoDeStringParaNumero()
	conversaoDeStringParaFoat()
}
