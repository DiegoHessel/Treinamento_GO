package main

import "fmt"

func fullName(firstName, lastName string){
	fmt.Printf("Nome Completo: %s %s\n", firstName, lastName )
}

func main(){
fmt.Println("Utilizando Função com parâmetros")
fullName("Diego", "Hessel")

}
