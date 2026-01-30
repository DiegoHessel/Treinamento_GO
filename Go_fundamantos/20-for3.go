package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Gerador de números aleatorios
	target := rand.Intn(10) + 1 // 1 a 10

	fmt.Println("Jogo da adivinhação")
	fmt.Println("Ente adivinhar um número entre 1 a 10")

	var guess int

	for {
		fmt.Println("Digite seu palpite ou 0 para sair")
		fmt.Scan(&guess) // Corrigido para usar fmt.Scan

		if guess == 0 {
			fmt.Println("Você desistiu. O numero era:", target)
			break
		}
		if guess > target {
			fmt.Println("Seu palpite é maior. Tente novamente")
		} else if guess < target {
			fmt.Println("Seu palpite é menor. Tente novamente")
		} else {
			fmt.Println("Parabéns, você acertou! O número era:", target)
			break
		}
	}
}
