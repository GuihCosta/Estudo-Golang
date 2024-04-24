package main

import "fmt"

func main (){

	const K float64 = 373.2

	C:= K-273

	fmt.Printf("O Ponto de ebulição da água em Kelvin (K) é %.1f K \nConvertendo para Celsius (°C) a temperatura é %.0f °C",K,C)

}