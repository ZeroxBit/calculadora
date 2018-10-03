package main

import (
	"fmt"
)

func main() {
	getArray()
	getSlice()
}

func getSlice() {
	var slice1 []string
	slice1 = append(slice1, "carlos", "maria")
	slice1 = append(slice1, "eduardo")
	fmt.Println(slice1)
}

func getArray() {
	var nombres [2]string
	numeros := [4]int{1, 2, 3, 4}
	nombres[0] = "carlos"
	nombres[1] = "maria"

	fmt.Println(nombres, numeros)
}
