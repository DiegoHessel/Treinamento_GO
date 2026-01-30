package main

// pacote para automatizar Tarefas no sistema operacional
//os = Operational System
import (
	"fmt"
	"os"
)

func main() {
	// Criando um arquivo
	arquivo, err := os.Create("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao Criar o arquivo:", err)
		return
	}
	defer arquivo.Close()

	// escrevendo no aequivo
	// usso o _ para não usar o primeiro retorno
	_, err = arquivo.WriteString("Texto de exemplo no arquivo")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}
	fmt.Println("Arquivo Criado e texto escrito com sucesso!")

	// lendo o arquivo
	conteudo, err := os.ReadFile("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao lerno arquivo:", err)
		return
	}
	fmt.Println("Conteudo do arquivo: ", string(conteudo))

}
