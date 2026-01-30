package main

import "fmt"

func main() {
	var nota float64

	fmt.Println("Informe a nota do aluno de 0 a 10")
	fmt.Scan(&nota)

	switch {
	case nota >= 9:
		fmt.Println("Classificação : A - Exemlente")
	case nota >= 7:
		fmt.Println("Classificação : B - Bom")
	case nota >= 5:
		fmt.Println("Classificação : C - Regular")
	case nota >= 3:
		fmt.Println("Classificação : D - Insuficiente")
	case nota >= 0:
		fmt.Println("Classificação : F - Reprovado")
	default:
		fmt.Println("Nota invalida! Digite uma nota entre 0 e 10")
	}

}
