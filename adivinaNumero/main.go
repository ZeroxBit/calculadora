package main

import (
	"fmt"
)

func main() {
	playGame()
}

func playGame() {
	fmt.Println("Bienvenido al juego de azar")
	var num int
	for {
		fmt.Println("ingresar solo numeros pares del 1 al 10")
		fmt.Scanf("%d", &num)
		if num < 1 || num > 10 {
			fmt.Printf("Error el numero %d, no esta en el rango del 1 al 10, Ingrese otro numero \n", num)
			continue
		} else if num%2 != 0 {
			fmt.Printf("Disculpe %d, no es un numero par, Ingrese otro numero \n", num)
			continue
		} else if num%2 == 0 {
			resp, cont := getWin(&num)
			if follow := conti(&resp, &cont); follow {
				continue
			} else {
				fmt.Printf(resp+"\n", num)
				break
			}
		}
	}
}

func conti(cadena *string, flag *bool) bool {
	var cont int
	if *flag {
		fmt.Println("numero incorrecto, desea seguir jugando? 1->si, 2->no")
		fmt.Scanf("%d", &cont)
		if cont == 1 {
			return true
		}
	}
	return false
}

func getWin(num *int) (string, bool) {
	var resp string
	if numWin := 2; numWin == *num {
		resp = "Felicitaciones el numero %d, fue el correcto"
		return resp, false
	} else {
		resp = "Lo sentimos, no es el numero correcto"
		return resp, true
	}
}
