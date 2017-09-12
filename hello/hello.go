package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	exibeMenu()

	switch leComando() {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1

	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O valor da variável comando é:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)
}
