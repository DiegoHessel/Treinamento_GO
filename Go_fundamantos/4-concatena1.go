package main

import (
	"fmt"
)

func main(){
	var nome string
	var idade int
	var altura float64
	var maiorDeIdade bool

	fmt.Println("Informe seu nome:")
	fmt.Scan(&nome)
	fmt.Println("Informe sua idade:")
	fmt.Scan(&idade)
	fmt.Println("Informe sua altura:")
	fmt.Scan(&altura)


	maiorDeIdade = idade >= 18
	fmt.Println("\n Dados Pessoais \n")
	fmt.Printf("Nome: %s",nome)
	fmt.Printf("Atura:%.2f",altura)
	fmt.Printf("Idade:%d", idade)
	fmt.Println("Maior de idade %v",maiorDeIdade)

		fmt.Printf("Meu nome é: %s\nMinha idade é: %d\nMinha altura é: %.2f\nSou maior de idade? %t\n", nome, idade, altura, maiorDeIdade)
	}
}