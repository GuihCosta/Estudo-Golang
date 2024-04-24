package main //1° Declarar pacote principal

import "fmt" //2° Importar função fmt

const ebulicaoF = 212.0 //Declarar valor constante

//3° Declarar função principal:
func main() {
	var F = ebulicaoF
	var C = (F - 32) * 5 / 9
	fmt.Println("Ponto de ebulição da água em °F é:", F)
	fmt.Println("Ponto de ebulição da água em °C é:", C)
}
