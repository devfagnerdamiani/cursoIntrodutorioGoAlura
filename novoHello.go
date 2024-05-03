package main

import "fmt"

func executar() {

	nome := "Fagner"
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("A versão deste programa é: ", versao)

	fmt.Println("O tipo de váriavel nome é: ", nome)
	fmt.Println("O tipo de váriavel versao é: ", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")

	var comando int

	fmt.Scanln("%d%", &comando)

}
