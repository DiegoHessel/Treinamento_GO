package main

import (
	"fmt"
	"strings"
)

func main() {
	// criando um map para contar as palavras

	text := "go é uma lingagem de programação e go é muito rápida"
	words := strings.Fields(text)

	wordCount := make(map[string]int)

	//Contagem da frequenncia de palavras
	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println("Contagem de Palavas")
	for word, count := range wordCount {
		fmt.Printf("Palavras %s | Frequência %d \n ", word, count)
	}
}
