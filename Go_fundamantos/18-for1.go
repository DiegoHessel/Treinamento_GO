package main

import "fmt"

func main(){
	nomes := [4]string{"Ana","Carlos","Beatriz","Joáo"}
	
	fmt.Println("Estrutura de Repetição For")
	 
	for id, nome := range nomes{
		//fmt.Println("ID", id)
		// fmt.Println("Nome", nome)
		fmt.Println(id, "->", nome)
		 if len(nome) > 5{
			fmt.Println(nome, " tem MAIS de 5 caracteres")
		 } else {
			fmt.Println(nome, " tem MENOS de 5 caracteres")
		}
	}
}
