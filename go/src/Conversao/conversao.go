package main

import "fmt"

var a int = 20
var b string = "11"

func main (){

	var c float64 = float64(a) // para converter um tipo em outro, utilizamos o tipo que desejamos e a variavel 
	fmt.Printf("O valor de C é:%.2f e o valor de B é: %s ",c,b)

}