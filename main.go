package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	name           = "Michael Scott"
	version        = 1.0
	xMonitoring    = 3
	timeMonitoring = 5 // em segundos
	sucessCode     = 200
	sitesFile      = "sites.txt"
	logFile        = "log.txt"
)

func main() {
	fmt.Println("Olá", name, ", bem vindo a versão ", version, "da aplicação")
	fmt.Println("")
	menu()
}

func menu() {
	for {
		var option int

		fmt.Println("Escolha uma opção:")
		fmt.Println("1 - Iniciar monitoramento")
		fmt.Println("2 - Logs")
		fmt.Println("3 - Encerrar programa")
		fmt.Scan(&option)
		fmt.Println("")
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
		fmt.Print("Encerrando programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando desconhecido, tente novamente")
		menu()
	}
}

func monitoring() {
	for x := 0; x < xMonitoring; x++ {
		fmt.Println("Iniciando monitoramento...")
		webSites := readFile()

		randomLink := "https://httpstat.us/" + randomNumber()
		webSites = append(webSites, randomLink)

		for _, site := range webSites {
			testWebSite(site)
		}
		fmt.Println("verificação", x+1, "concluída")
		fmt.Println("")
		if x != xMonitoring-1 {
			time.Sleep(timeMonitoring * time.Second)
		}
	}
}

func testWebSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("#### Ocorreu um erro ao verificar o site ", site, "---- ", err)
		os.Exit(-1)
	}
	if resp.StatusCode == sucessCode {
		fmt.Println("Site: ", site, " carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site: ", site, " esta com problema! Code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func log() {
	fmt.Println("Exibindo logs...")
	file, err := os.ReadFile(logFile)

	if errors.Is(err, fs.ErrNotExist) {
		fmt.Println("\nNenhum Log foi criado, inicie o monitoramento!")
	} else if err != nil {
		fmt.Println("Erro inesperado: ", err)
	}
	fmt.Println(string(file))
}

func randomNumber() string {
	array := []string{"100", "101", "102", "103", "200", "201", "202", "203", "204", "205", "206", "207", "208", "226", "300", "301", "302", "303", "304", "305", "306", "307", "308", "400", "401", "402", "403", "404", "405", "406", "407", "408", "409", "410", "411", "412", "413", "414", "415", "416", "417", "418", "421", "422", "423", "424", "425", "426", "428", "429", "431", "451", "500", "501", "502", "503", "504", "505", "506", "507", "508", "510", "511", "419", "420", "440", "444", "449", "450", "460", "463", "494", "495", "496", "497", "498", "499", "520", "521", "522", "523", "524", "525", "526", "527", "530", "561"}

	randomIndice := rand.Intn(len(array))

	return array[randomIndice]
}

func readFile() []string {
	var response []string
	file, err := os.Open(sitesFile)

	if err != nil {
		fmt.Println("#### Ocorreu um erro -", err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		response = append(response, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return response
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro em registerLog: ", err)
		os.Exit(-1)
	}

	now := time.Now().Format("02/01/2006 15:04:05")
	file.WriteString("#### " + now + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}
