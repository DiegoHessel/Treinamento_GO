package main

import (
	"fmt"
	"strings"
)

func main() {
	movieName := "Top Gun"
	movieName2 := "top Gun"

	fmt.Println(movieName == movieName2)
	movieDescripition := `
Top Gan Maveric é um filme de aviação e 
aventura muito conseituado 
na indústria
`
	// conversão de de palavras para MAÍUSCULO
	fmt.Println(strings.ToUpper(movieDescripition))

	// conversão de de palavras para minúsculo
	fmt.Println(strings.ToLower(movieDescripition))

	// Primeira letra em maiúsculo
	fmt.Println(strings.Title(movieDescripition))

	// Encontrar a posição de um caractere
	fmt.Println(strings.Index(movieDescripition, "p"))

	// contando o numero de ocorrência de um caractere
	fmt.Println(strings.Count(movieDescripition, "a"))
	fmt.Println(strings.Count(movieDescripition, "e"))

	// substitua um elemento por outro
	fmt.Println(strings.ReplaceAll(movieDescripition, "filme", "serie"))
}
