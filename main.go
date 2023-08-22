package main

import "fmt"

func main() {
	nome := "Michael Scott"
	versao := 0.1
	fmt.Println("Olá", nome, ", bem vindo a versão ", versao, "da aplicação")

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Logs")
	fmt.Println("0 - Encerrar programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println(comando)
}
