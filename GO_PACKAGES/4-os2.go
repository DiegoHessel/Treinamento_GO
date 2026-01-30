package main

// pacote para automatizar Tarefas no sistema operacional
//os = Operational System
import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// . 1. retorna a pasta atual
func getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório: ", err)
		return ""
	}
	return dir
}

// 2- listar arquivos e pastas
func listFelesAndDirectories() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Erro ao listar arquivos: ", err)
		return
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// 3- Versáo do SO

func getOsVersion() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "ver")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("uname", "-a")
	} else if runtime.GOOS == "darwin" {
		cmd = exec.Command("sw_vers")
	} else {
		fmt.Println("SO não Suportado!")
		return
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao executar comando: ", err)
		return
	}
	fmt.Println(string(out))
}

// 4- configuração da maquina

func getSistemInfo() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "systeminfo")
	} else {
		fmt.Println("SO não suportado para obter informações do sistema!")
		return
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao executar comando: ", err)
		return
	}
	fmt.Println(string(out))
}

// 5- desligar o computador em 1h
func shutdownInOneHour() {
	cmd := exec.Command("shutdown", "/s", "/t", "3600")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao agendar desligamento")
	}
}
//6- cancelar o desligamento 
func cancelShutdown() {
	cmd := exec.Command("shutdown", "/a")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao agendar desligamento")
	}
}

func main() {
	fmt.Println("Pasta Atual:", getCurrentDirectory())
	fmt.Println("Arquivos e Pasta:")
	listFelesAndDirectories()

	fmt.Println("Versão do SO:")
	getOsVersion()

	fmt.Println("Configuração da Máquina:")
	getSistemInfo()
	//shutdownInOneHour()
	cancelShutdown()
}
