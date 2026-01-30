package main

import "fmt"

func main() {
	// criando um map como nome do aluno como chave e a nota como valor
	estudantes := map[string]float64{
		"Ana":     7.5,
		"Carlos":  4.3,
		"Beatriz": 8.9,
		"Joao":    6.2,
	}
	fmt.Println(estudantes)

	fmt.Println("Classificação dos Alunos:")

	for nome, nota := range estudantes {
		status := "Reprovado"

		if nota >= 6.0 {
			status = "Aprovado"
		}
		fmt.Printf("Aluno: %s | Nota: %.2f | Status: %s \n", nome, nota, status)
	}
}
