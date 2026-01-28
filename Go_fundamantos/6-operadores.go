package main
import "fmt"

func main(){
	var num1, num2 int
	fmt.Println("Informe o seu primeiro numero")
	fmt.Scan(&num1)

	fmt.Println("Informe seu segundo numero")
	fmt.Scan(&num2)
	//Aritiméticos
	soma := num1 + num2
	sub :=  num1 - num2
	mult := num1 * num2
	div :=  num1 / num2
	mod := num1 % num2


	fmt.Printf("Soma de %d e %d é %d \n",num1,num2,soma)
	fmt.Printf("Subitração de %d e %d é %d \n",num1, num2,sub)
	fmt.Printf("Multiplicação de %d e %d é %d \n",num1,num2,mult)
	fmt.Printf("Divisão de %d e %d é %d\n",num1,num2,div)
	fmt.Printf("Divisão de %d e %d é %d\n",num1,num2,div)
	fmt.Printf("Resto da Divisão de %d e %d é %d\n",num1,num2,mod)
	
	//Atribuição
	num1 +=1 // num1 = num1 + 1
	num1 -=1 // num1 = num1 - 1
	num1 *=1 // num1 = num1 * 1
	num1 /=1 // num1 = num1 / 1
	
	println("Valor final de num1 após alteração de atribuição: ",num1)

	//Comparação 
	bigger := num1 > num2
	smaller := num1 < num2
	equal := num1 == num2
	different:= num1 != num2
	biggerEqual := num1 >= num2
	smallerEqual := num1 <= num2

	fmt.Printf("O número %d é maior que %d? %v \n",num1,num2,bigger)
	fmt.Printf("O número %d é menor que %d? %v \n",num1,num2,smaller)
	fmt.Printf("Os números %d e %d sáo iguais? %v \n",num1,num2,equal)
	fmt.Printf("Os números %d e %d são diferentes? %v \n",num1,num2,different)
	fmt.Printf("O número %d é maior ou igual a %d? %v \n",num1,num2,biggerEqual)
	fmt.Printf("O número %d é menor ou igual a %d? %v \n",num1,num2,smallerEqual)

}

