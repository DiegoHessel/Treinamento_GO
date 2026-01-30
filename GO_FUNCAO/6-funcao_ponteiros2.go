package main

import "fmt"

// função que recebe um ponteiro como argumento 
// e altera o valor da variavel original
type Cliente struct{
	Nome string
	Idade int
}
// Função para receber um ponteiro para alterar os dados do cliente 

func atualizarCliente(c *Cliente, novoNome string, novaIdade int){
	c.Nome = novoNome
	c.Idade = novaIdade
}
func  main(){
cliente := Cliente{Nome: "Diego", Idade : 23}
fmt.Println("Valor inicial: \n",cliente)
fmt.Printf("Valor inicial Nome: %s Idade: %d \n",cliente.Nome,cliente.Idade)
	// passando o ponteriro para a função
	atualizarCliente(&cliente,"Diego Hessel",22)
	fmt.Println("Valor após a alteração: \n",cliente)
	fmt.Printf("Valor após a alteração Nome: %s Idade: %d \n",cliente.Nome,cliente.Idade)
}