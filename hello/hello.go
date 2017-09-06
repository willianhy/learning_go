package main

import "fmt"

func main(){
    nome := "Douglas"
    versao := 1.1

    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)

    fmt.Println("O valor da variável comando é:", comando)

    switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo logs...")
    case 0:
        fmt.Println("Saindo do programa")
    default:
        fmt.Println("Não conheço este comando")
    }
}