package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {

	exibirIntroducao()
	leSitesDoArquivo()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}

}

func nomeEidade() (string, int) {
	nome := "Thallys"
	idade := 20
	return nome, idade
}

func exibirIntroducao() {
	nome := "Thallys"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("")

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i)
			testandoSite(site)

		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testandoSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt") //apenas abrir o arquivo com o endereço de memória

	//arquivo, err := ioutil.ReadFile("sites.txt") //ler todo arquivo de uma vez

	if err != nil { //identificando o erro e sinalizando ele por conta do Open devolver dois valores
		fmt.Println("Ocorreu um erro!", err)
	}

	leitor := bufio.NewReader(arquivo) // ler o arquivo em uma linha

	for {
		linha, err := leitor.ReadString('\n') //ler uma linha e retornar uma string
		//'\n' serve para delimitar até onde ele tem que ler, nesse caso até a quebra da linha
		linha = strings.TrimSpace(linha) //tirar espaços das linhas e "\n" tambem
		sites = append(sites, linha)     //colocar linhas dentro do array de sites

		if err == io.EOF { //identificando o erro e sinalizando ele por conta do Open devolver dois valores
			break //Quando ele  encontra o erro end of file ele saira do loop
		}

	}
	arquivo.Close() //finalizar arquivos para liberar pra outras pessoas usarem
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//cria o arquivo log.txt e Lê passando a permissao do windows

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "_ online: " + strconv.FormatBool(status) + "\n") // funcao para escrever uma string site e o format bool no log.txt
	// formato de hora fica disponível na documentação
	arquivo.Close()
	//fechando arquivo
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt") //ler arquivo todo de uma vez

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
