package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	menu()
}

func menu() {
	for {
		name := "Michael Scott"
		var version float32 = 0.2
		var option int

		fmt.Println("Olá", name, ", bem vindo a versão ", version, "da aplicação")

		fmt.Println("1 - Iniciar monitoramento")
		fmt.Println("2 - Logs")
		fmt.Println("3 - Encerrar programa")
		fmt.Println("Escolha uma opção:")
		fmt.Scan(&option)

		selection(option)

	}
}

func selection(option int) {

	switch option {
	case 1:
		monitoring()
	case 2:
		log()
	case 3:
		turnOff()
	default:
		fmt.Println("Comando desconhecido, tente novamente")
		menu()
	}
}

func monitoring() {
	fmt.Println("Iniciando monitoramento...")
	webSite := "http://www.alura.com.br"
	// https://httpstat.us/ sugestão de site para testes
	resp, _ := http.Get(webSite)
	if resp.StatusCode == 200 {
		fmt.Print("Site: ", webSite, " carregado com sucesso!")
	} else {
		fmt.Println("Site: ", webSite, " esta com problema! Code:", resp.StatusCode)
	}
}

func log() {
	fmt.Println("Exibindo logs...")
}

func turnOff() {
	fmt.Println("Encerrando programa...")
	i := 0
	for i < 50 {
		i++
		fmt.Println(".")
	}
	os.Exit(0)
}
