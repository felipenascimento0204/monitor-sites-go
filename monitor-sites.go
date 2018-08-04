package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"io"
	"fmt"
	"os"
	"net/http"
	"time"
	"bufio"
)

const monitoramento = 4
const versao float32 = 1.1

func main() {

	fmt.Println("Monitor de sites versão: ", versao)

	for{

		exibeMenu()

		comando := leComando()

		switch comando {
			case 1: 
				iniciarMonitoramento()
			case 2:
				imprimeLogs()
			case 0:
				sair()
			default:
				comandoDesconhecido()
		}

	}

}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func exibeMenu(){
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair do programa")
}

func sair(){
	fmt.Println("Saindo do programa")
	os.Exit(0)
}

func comandoDesconhecido(){
	fmt.Println("Comando desconhecido")
	os.Exit(-1)
}

func iniciarMonitoramento(){
	
	fmt.Println("Iniciando monitoramento")
	sites := leSitesArquivo()

	fmt.Println("Quantidade de sites:", len(sites))

	fmt.Println("")
	for i := 0; i < 5; i++{
		for i, site := range sites{
			fmt.Println("Testando site", i + 1)
			verificaHost(site)
		}
		time.Sleep(monitoramento * time.Second)
		fmt.Println("")
	}

}

func verificaHost(site string){
	
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro ao verificar o siste:", site, "Erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(site, "SUCESSO")
		registraLog(site, true)
	} else {
		fmt.Println(site, "ERRO Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesArquivo() []string{

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo:", err)
		sair()
	}

	leitor := bufio.NewReader(arquivo)

	for{
		linha, err := leitor.ReadString('\n')
		if err == io.EOF {
			break
		}
		sites = append(sites, strings.TrimSpace(linha))
	
	}
	return sites
}
		
func registraLog(site string, status bool){
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs(){
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro ao ler os arquivos de log", err)
	}
	fmt.Println(string(arquivo))
}