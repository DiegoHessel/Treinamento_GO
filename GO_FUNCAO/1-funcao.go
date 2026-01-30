package main

import "fmt"

// 1- imprimndo uma mensage de vias vindas

func welcome() {
	fmt.Println("Bem vindos ao sistema de filme")
}

// 2 - funcao de cadastro de filmes
func createMovie() {
	var nome string
	var yearRelease int
	var moviePrice float64

	fmt.Println("Digite o nome do filme:")
	fmt.Scan(&nome)
	fmt.Println("Digite o Ano de lançamento:")
	fmt.Scan(&yearRelease)
	fmt.Println("Digite o Preço do filme:")
	fmt.Scan(&moviePrice)

	fmt.Printf("Nome do fime: %s Ano de Lançamento(%d) Preço do Filme- R$ %.2f \n", nome, yearRelease, moviePrice)
}

// calcular media de notas
func calculateAverage() float64 {
	var numRatings int
	fmt.Println("Digite quantas avaliações deseja fazer para o filme:")
	fmt.Scan(&numRatings)

	var total float64
	for i := 0; i < numRatings; i++ {
		var nota float64
		fmt.Println("Digite a nota para o filme:")
		fmt.Scan(&nota)
		total += nota
	}

	var average float64
	if numRatings > 0 {
		average = total / float64(numRatings)
	} else {
		average = 0
	}
	return average
}

func main() {
	fmt.Println("Utilizando função")
	welcome()
	createMovie()
	media := calculateAverage()
	fmt.Printf("A média das avaliações é: %.2f \n", media)
}
