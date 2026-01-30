package main

import "fmt"

// função para soma de numeros com variadicas
func sum(nums ...int) {
	sumTotal := 0
	for _, n := range nums {
		sumTotal += n
	}
	fmt.Printf("Soma é: %d \n", sumTotal)
}

//2- função para apresentação de cursos com variadicos

func presentation(data map[string]string) {
	for kay, value := range data {
		fmt.Printf("%s - %s\n", kay, value)
	}
}

func main() {
	sum(7)
	sum(7, 8)
	sum(7, 8, 9)

	presentation(map[string]string{
		"name":     "Python",
		"cadegory": "Backend",
		"level":    "Iniciante",
	})

	presentation(map[string]string{
		"name":     "Visão Computacional com python",
		"cadegory": "IA",
		"level":    "Avançado",
	})

	presentation(map[string]string{
		"name":     "Dashboard",
		"cadegory": "Data Science",
		"level":    "Intermediario",
	})
}
