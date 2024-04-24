package main //1° Declarar pacote principal

import "fmt" //2° Importar função fmt

const ebulicao1F = 212.0 //Declarar valor constante

//3° Declarar função principal:
func main() {
	F := ebulicao1F // := declara e atribui o valor direto a variavel, sem necessitar declara-la antes 
	C := (F - 32) * 5 / 9
	fmt.Printf("Ponto de ebulição da água em °F é = %.2f, %T \nPonto de ebulição da água em °C é: = %.2f,%T ",F,F,C, C) // PrintF pode ser usado para colocar varias variaveis numa linha só e %T mostra o Tipo da variavel utilizada.
	
}
