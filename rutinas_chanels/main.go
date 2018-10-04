package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////////////////
//para mandar informacion a los chanes se hace de esta manera Ej:
// -- a := make(chan string)
// -- a <- "hola"  // aqui se esta mandando "hola" al chanel "a"
//para recibir info se tiene que esta del lado izquierdo de la flecha y
//para mandar tiene que estar del lado derecho.
//////////////////////////////////////////////////////////////////////////

//de la manera que esta declarado chan<- quiere decir que va ha ser solo de escritura, no se va a poder leer
func ping(ball chan<- int, action chan<- string) {
	ball <- 1
	action <- "player ping"
}

func pong(ball chan<- int, action chan<- string) {
	ball <- 2
	action <- "player pong"
}

//aqui el chanes (<-chan) es de solo lectura, no se puede reescribir
func referee(action <-chan string) {
	for {
		fmt.Println("Action: ", <-action)
	}
}

func pingPong() {
	ball := make(chan int)
	action := make(chan string)
	go referee(action)
	go ping(ball, action)
	for {
		value := <-ball
		switch value {
		case 1:
			go pong(ball, action)

		case 2:
			go ping(ball, action)
		}
	}
}

func main() {
	go pingPong()
	time.Sleep(10 * time.Second)
}
