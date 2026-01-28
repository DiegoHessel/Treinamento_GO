package main

import ("fmt"
		"strings"
)

func main(){
	movieName := "Top Gun"
	movieName2 := "top Gun"

	fmt.Println(movieName == movieName2)
	movieDescripition := `
Top Gan Maveric é um filme de aviação e 
aventura muito conseituado 
na indústria
`
	line := "="
	//quantas vezes vc quer repetir algum caracter 
	fmt.Println(strings.Repeat(line,40))


	//fmt.Printf("========================")
	fmt.Println(movieDescripition)

	//verifica se uma palavra existe dentro de uma string
	fmt.Println(strings.Contains(movieDescripition,"top"))
	fmt.Println(strings.Contains(movieDescripition,"filme"))
}