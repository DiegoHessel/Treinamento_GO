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
	fmt.Println(nome)
	fmt.Println(altura)
	fmt.Println(idade)


	maiorDeIdade = idade >= 18
	fmt.Println(nome)
	fmt.Println(altura)
	fmt.Println(idade)
	fmt.Println(maiorDeIdade)
}