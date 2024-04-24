package main

import "fmt"

func main(){
	/*Tipos: Numericos (int,float)*/
	fmt.Println("2.5 + 3.5=", 2.5 + 3.5)

	/*Tipos: Strings*/
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[2])
	fmt.Println("Hello" + " World")

	/*Tipos: Booleanos (&& = e, ||= ou, != negação)  */
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!false)
	fmt.Println(!true)

	/*Inferencia de Tipos */
}