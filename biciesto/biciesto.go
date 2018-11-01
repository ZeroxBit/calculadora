package main

import "fmt"

func main() {
	var anio uint
	fmt.Println("ingrese el año")
	fmt.Scanf("%v", &anio)
	if anio%4 == 0 && (anio%100 != 0 || anio%400 == 0) {
		fmt.Printf("el año %v es par ", anio)
	} else {
		fmt.Printf("el año %v, es impar", anio)
	}
}
