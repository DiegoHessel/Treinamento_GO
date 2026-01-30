package main

import "fmt"

//struct estrutura de modelo parecido com classe em Java ou C ou C++
// Definindo a struct Carro
type Pessoa struct{
	Nome string
	Idade int
	Endereco string

}
func main(){
// criar a instância da struct carro
pessoa := Pessoa{
	Nome : "Diego Hessel",
	Idade : 23,
	Endereco :"Rua Tutoia",
}
	fmt.Println("Informacões da Pessoa")
	fmt.Printf("Nome: %s \n", pessoa.Nome)
	fmt.Printf("Idade: %d \n", pessoa.Idade)
	fmt.Printf("Endereço: %s \n", pessoa.Endereco)
}

