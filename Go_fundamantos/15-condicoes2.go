package main
import "fmt"

func main(){
	var nota float64

	fmt.Println("Digite sua nota de 0 a 10")
	fmt.Scan(&nota)
	// condoicoies 	
	// e = &  
	// ou = ||
	if nota < 0 ||  nota > 10{
		fmt.Println("Nota inválida! Digite uma nota entre 0 e 10")
	} else if nota >=9 {
		fmt.Println("Exelente! Aprovado com Distinção")
	}else if nota >=7 {
		fmt.Println("Muito bom! Aprovado")
	} else if nota >=5 {
		fmt.Println("Satisfatório! Aprovado")
	}else{
		fmt.Println("Reprovado")
	}
}