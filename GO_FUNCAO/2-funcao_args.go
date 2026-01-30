package main

import "fmt"

func fullName(firstName, lastName string) {
	fmt.Printf("Nome Completo: %s %s\n", firstName, lastName)
}
func sumNumbers(a, b int) int {
	return a + b
}

func address(country string) {
	if country == "" {
		country = "Brasil"
	}
	fmt.Printf("Eu moro no %s \n", country)
}
func main() {
	fmt.Println("Utilizando Função com parâmetros")
	fullName("Diego", "Hessel")
	fmt.Printf("Soma: %d \n", sumNumbers(80, 100))
	address("")
	address("Chile")
}
