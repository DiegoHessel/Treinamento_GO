package main

import "fmt"

func main() {
	//tipo de dado inteito usar int
	var idade int = 30
	// dados de ponto flutuante usar float
	var altura float64 = 1.75
	//Tipo de dado booleando (verdadeiro ou falso)
	var mairDeIdade bool = idade >= 18
	// Tipo String
	var nome string = "Diego"
	fmt.Println("Dados pessoais")
	fmt.Println("Nome:")
	fmt.Println(nome)
	fmt.Println("IDADE:")
	fmt.Println(idade)
	fmt.Println("Altura :")
	fmt.Println(altura)
	fmt.Println("maior de idade :")
	fmt.Println(mairDeIdade)
	fmt.Println(fmt.Sprintf("%T", mairDeIdade))

	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("IDADE: %d\n", idade)
	fmt.Printf("Altura: %.2f\n", altura)
	fmt.Printf("Maior de idade: %t\n", mairDeIdade)
	fmt.Printf("Tipo de dado de 'mairDeIdade': %T\n", mairDeIdade)
}
