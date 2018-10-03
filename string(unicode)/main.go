package main

import "fmt"

func main() {
	// こんにちは世界 = hola mundo (en japones)
	fmt.Println("こんにちは世界"[3])         // esto regresa el valor ascii de la posicion 3 de la cadena
	fmt.Println(string("こんにちは世界"[3])) // para obtener el valor del indice
	imprimir()
}

func imprimir() { // para imprimir los caracteres individuales
	str := "こんにちは世界"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}
}
