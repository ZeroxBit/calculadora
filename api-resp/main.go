package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) { // el primer parametro es donde se escribiera la respuesta y el segundo es la respuesta o datos
	fmt.Fprintf(w, " Hola Mundo ")
}

func prueba(w http.ResponseWriter, r *http.Request) { // el primer parametro es donde se escribiera la respuesta y el segundo es la respuesta o datos
	fmt.Fprintf(w, "Hola Mundo desde prueba ")
}

func usuario(w http.ResponseWriter, r *http.Request) { // el primer parametro es donde se escribiera la respuesta y el segundo es la respuesta o datos
	fmt.Fprintf(w, "Hola usuario")
}

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func main() {

	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/", home)

	serverMux.HandleFunc("/prueba", prueba)

	serverMux.HandleFunc("/usuario", usuario)

	http.ListenAndServe(":8080", serverMux)
}
