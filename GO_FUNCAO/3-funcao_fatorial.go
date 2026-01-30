// 2 -> 2 * 1
// 3 -> 3 * 2 * 1

package main

import "fmt"

func factorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

func totalSum(num int) int {
	if num == 1 {
		return 1
	} else {
		return num + totalSum(num-1)
	}
}
func main() {
	var number int
	var num int

	fmt.Printf("Digite o número para o fatorial:")
	fmt.Scan(&number)
	fmt.Printf("O fatorial de %d é %d \n", number, factorial(number))

	fmt.Println("Digite um numero para soma:")
	fmt.Scan(&num)
	fmt.Printf("A soma total do %d é %d \n", num, totalSum(num))
}
