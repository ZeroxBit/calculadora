package main

import (
	"fmt"
)

func main() {
	getArray()
	getSlice()
	showIndexValue()
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

func showIndexValue() {
	x := []int{1, 2, 3, 4, 5, 6, 7}

	for i, v := range x {
		fmt.Println("index: ", i, " ", "value: ", v)
	}
}
