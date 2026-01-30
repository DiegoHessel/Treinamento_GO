package main

import "fmt"

// struct estrutura de modelo parecido com classe em Java ou C ou C++
// Definindo a struct Carro
type Carro struct {
	Modelo string
	Ano    int
	Cor    string
}

func main() {
	// criar a instância da struct carro
	carro1 := Carro{
		Modelo: "Fusca",
		Ano:    1965,
		Cor:    "Azul",
	}
	fmt.Println("Informacões do Carro")
	fmt.Printf("Modelo do Carro: %s \n", carro1.Modelo)
	fmt.Printf("Ano do Carro: %d \n", carro1.Ano)
	fmt.Printf("Cor do Carro: %s \n", carro1.Cor)

}
