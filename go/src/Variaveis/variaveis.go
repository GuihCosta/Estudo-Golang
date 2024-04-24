package main

import "fmt"

/*Podemos declarar as variáveis fora do scopo da função main*/
var nome string = "Guilherme"
var idade int = 23
var versao float32 = 3.2

var nome2 = "Guilherme"
var idade2 = 23
var versao2 = 3.2

var x string = "Hello, World!"

func main (){
	/*Declarando o tipo da variável*/
	fmt.Println("Meu nome é ",nome, " e minha idade é ", idade, " anos")
	fmt.Println("Estou usando o programa de versão: ", versao)

	/*Sem declarar o tipo da variável (Não é necessário declarar o próprio go sabe o tipo da variável)*/
	fmt.Println("\nMeu nome é ",nome2, " e minha idade é ", idade2, " anos")
	fmt.Println("Estou usando o programa de versão: ", versao2)

	fmt.Print("\n",x)

}