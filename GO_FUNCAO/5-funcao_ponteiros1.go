package main

import "fmt"

// função que recebe um ponteiro como argumento 
// e altera o valor da variavel original
func  alterValue(x *int){

	//Alterando o valor do ponteiro
	*x = *x * 2
}

func  main(){
	num := 5
	fmt.Printf("Valor inicial: %d \n",num)
	// passando o ponteriro para a função
	alterValue(&num)
	fmt.Printf("Valor após a alteração: %d \n",num)

}