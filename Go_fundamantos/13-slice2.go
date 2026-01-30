package main

import "fmt"

func main() {
	//criando slice vazio de inteiros
	frutas := []string{"Banana", "Laranja", "Uva", "Maçã", "Morango"}

	//criando um sub slice
	subslice := frutas[1:4]

	fmt.Println("Slice de frutas:", frutas)
	fmt.Println("Subslice de frutas:", subslice)

	subslice[0] = "manga"
	fmt.Println("Slice de frutas após a modificação:", frutas)
}
