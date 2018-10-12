package main

import (
	"fmt"
)

func main() {
	// addSlice()
	// deleteSlice()
	sliceMultiDimen()
}

func sliceMultiDimen() {
	cl := []string{"carlos", "sanchez", "29"}
	mr := []string{"maria", "perez", "25"}

	arrPersonas := [][]string{cl, mr}

	fmt.Println(arrPersonas)
}

func addSlice() {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	z := []int{22, 44, 55}
	l := x[1:6] // para dividir un slice

	//agregando un slice a otro
	x = append(x, z...) // (z...) divide el slice y lo pasa uno por uno
	fmt.Println(l)
	fmt.Println(x[2:5])

	for i, v := range x {
		fmt.Println("index: ", i, "value: ", v)
	}
}

func deleteSlice() {
	x := []int{1, 2, 3, 4, 5, 6, 75, 44, 22, 88}
	fmt.Println(x)
	x = append(x[:2], x[7:]...) // quitando el 4,5,6 del slice
	fmt.Println(x)
}
