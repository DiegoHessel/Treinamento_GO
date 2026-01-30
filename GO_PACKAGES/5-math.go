package main

import (
	"fmt"
	"math"
	"strings"
)

// Acessar o número de PI
func acessarPi() {
	fmt.Printf("PI Arredondado (2 casas decimais): %.2f \n", math.Pi)
}

// Acessar o número de PI
func acessarEuler() {
	fmt.Printf("Euler Arredondado (2 casas decimais): %.2f \n", math.E)
}

// 3 arredondando numeros para cima ou para baixo
func arredondar() {
	num := 10.4
	fmt.Println("Arredondando para cima:", math.Ceil(num))
	fmt.Println("Arredondando para baixo:", math.Floor(num))
}

// Funções para que o usuário insira os valores
func calcularPotencia() {
	var base, expoente float64
	fmt.Print("Potência \n")
	fmt.Print("Digite a base: ")
	fmt.Scan(&base)
	fmt.Print("Digite o expoente: ")
	fmt.Scan(&expoente)
	fmt.Printf("Potência de %.2f elevado a %.2f: %.2f\n", base, expoente, math.Pow(base, expoente))
}

func calcularRaizQuadrada() {
	fmt.Print("Raiz Quadrada \n")
	var numero float64
	fmt.Print("Digite o número para calcular a raiz quadrada: ")
	fmt.Scan(&numero)
	fmt.Printf("Raiz quadrada de %.2f: %.2f\n", numero, math.Sqrt(numero))
}

func calcularLogaritmo() {
	fmt.Print("Logaritmo \n")
	var numero float64
	fmt.Print("Digite o número para calcular o logaritmo: ")
	fmt.Scan(&numero)
	fmt.Printf("Logaritmo de %.2f: %.2f\n", numero, math.Log(numero))
}
func linhas() {
	line := "="
	//quantas vezes vc quer repetir algum caracter
	fmt.Println(strings.Repeat(line, 40))
}

func main() {
	acessarPi()
	acessarEuler()
	fmt.Println("Potência de 5 elevado 5 é:", math.Pow(5, 5))
	linhas()

	fmt.Println("Raiz quadrada de 169:", math.Sqrt(169))
	linhas()

	fmt.Println("Logaritimo de 10 =", math.Log(10))
	linhas()

	calcularRaizQuadrada()
	linhas()

	calcularLogaritmo()
	linhas()

	calcularPotencia()

}
