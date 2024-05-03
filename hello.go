package main

import (
	"bufio"     // Pacote de Leitor de arquivo linha a linha
	"fmt"       // Pacote de imprimir na tela
	"io"        // Pacote de abrir arquivo
	"io/ioutil" // Pacote para Escrever no arquivo
	"net/http"  // Pacote de Testar site
	"os"        // Pacote de comunicação com sistema operacional
	"strconv"   // Pacote de conversão de String para Booleano
	"strings"   // Pacote de conversão para String
	"time"      // Pacote de manipulação de tempo
)

const monitoramentos = 2
const delay = 5

func main() {

	// fmt.Println("Monitorando...")
	// site com URL inexistente
	// var sites [4]string
	// sites[0] = "https://httpbin.org/status/404"
	// sites[1] = "https://alura.com.br"
	// sites[2] = "https://caelum.org"

	// fmt.Println(sites)
	// fmt.Println(reflect.TypeOf(sites))
	// exibeNomes()

	exibeIntroducao() // Função que não retorna nada e não envia nada

	//	leSitesDoArquivo()

	for {
		exibeMenu() // Função para exibir menu, esta função não retorna nada

		//	nome, idade := devolveMeuNomeEIdade() // variaveis recebem os dois retornos
		//	nome, _ := devolveMeuNomeEIdade() // somente a variavel nome recebe o retorno ( O " _ " serve para ignorar o valor retornado de algum parametro/ variavel )
		//	fmt.Println(nome, " | ", idade)

		comando := leComando()

		switch comando {

		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs ...")
			imprimeLogs()
		case 0:
			fmt.Print("Saindo do Programa ...")
			os.Exit(0) // Saida do programa com sucesso
		default:
			fmt.Println("Não conheço esse comando!")
			os.Exit(-1) // Saida do programa com erro inesperado
		}

	}
}

// Exibe parte do enunciado
func exibeIntroducao() {

	nome := "Fagner"
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("A versão deste programa é: ", versao)

	fmt.Println("O tipo de váriavel nome é: ", nome)
	fmt.Println("O tipo de váriavel versao é: ", versao)
}

// Faz leitura da resposta do usuário
func leComando() int {

	var comandoLido int

	fmt.Scan(&comandoLido)
	fmt.Scanf("%d%", &comandoLido)
	fmt.Println("O endereço da variável comando é : ", &comandoLido)
	fmt.Println("O comando escolhido foi : ", comandoLido)

	return comandoLido
}

// Exibe opções de seleção para o usuário
func exibeMenu() {

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

// Faz a verificação e a comunicação de páginas que serão monitoradas
func iniciaMonitoramento() {
	fmt.Println("Monitorando ...")

	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)
}

// func devolveMeuNomeEIdade() (string, int) {
// 	nome := "Fagner"
// 	idade := 35
// 	return nome, idade
// }

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// site com URL inexistente
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("######################## INICIO ###########################")
			fmt.Println("Testando site ", i, " : ", site)
			testaSite(site)
			fmt.Println("######################### FIM #############################")
			fmt.Println("")
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

// func exibeNomes() {
// 	nomes := []string{"FAGNER", "KAREN", "BRUNO", "GABRIELE"}
// 	fmt.Println("O meu slice tem ", len(nomes), " nomes.")
// 	fmt.Println("O meu slice tem a capacidade de armazenar ", cap(nomes), " nomes.")
// 	nomes = append(nomes, "ICHIGO")
// 	fmt.Println(nomes)
// 	fmt.Println(reflect.TypeOf(nomes))
// 	fmt.Println("O meu slice tem ", len(nomes), " nomes.")
// 	fmt.Println("O meu slice tem a capacidade de armazenar ", cap(nomes), " nomes.")
// }

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um Erro ao testar o site ====>  ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("Sites.txt") // nesta linha o algoritimo busca o arquivo dentro da pasta deste projeto
	//	arquivo, err := os.Open("c://x//Sites.txt") // nesta linha o algoritimo busca o arquivo na pasta raiz c://
	// arquivo, err := os.Open("c://Sites.txt") // nesta linha o algoritimo busca o arquivo insistente na pasta apontada
	//arquivo, err := ioutil.ReadFile("Sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um Erro ao abrir o arquivo ====>  ", err)
	}

	//	fmt.Println(arquivo)
	//fmt.Println(string(arquivo))
	leitor := bufio.NewReader(arquivo) // Leitor de linha a linha

	if err != nil {
		fmt.Println("Ocorreu um Erro ao ler/percorrer o arquivo ====>  ", err)
	}

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		//		fmt.Println("FOI ========> ", linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	fmt.Println(sites)

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {

		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " | " + site + " -  online : " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
