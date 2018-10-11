package main

import (
	"fmt"
	"strconv"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hola mundo desde iris</h1>")
	})

	app.Get("/json", func(ctx iris.Context) {
		resp := map[string]int8{"uno": 1, "dos": 2, "tres": 3}
		ctx.JSON(resp)
	})

	app.Get("/down", func(ctx iris.Context) { // descarga un archivo
		ctx.SendFile("./files/screen.png", "archivoDes.png") // primer parametro el archivo a descargar, el segundo el nombre y extension del archivo cuando se descargue
	})

	app.Get("/param/{params:int}", func(ctx iris.Context) { // se puede especificar el tipo de data a recibir
		//paramt := ctx.Params().Get("params")  -> para obtener la data que se manda al api, es con Get("param") y usando entre comillas "", string
		paramt, _ := ctx.Params().GetInt("params") // para ibtener el tipo de dato int
		fmt.Println(paramt)
	})

	app.Get("/user/{id regexp([0-9][A-Z])}", func(ctx iris.Context) { // para definir los tipos de caracteres permitidos y la cantidad permitida
		id := ctx.Params().Get("id")
		fmt.Println(id)
	})

	// funcion que se ejecuta al llamar al metodo get/id/{id}
	app.Macros().Int.RegisterFunc("isPair", func() func(string) bool {
		return func(paramValue string) bool {
			eval, _ := strconv.Atoi(paramValue)
			return eval%2 == 0
		}
	})

	// funcion que evalua si la ruta llamda no existe y si no es asi, se manda un mensaje especificado
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<h1>Ruta no encontrada. 404</h1>")
	})

	// isPair() llama a la funcion una ves es llamado este end-point
	app.Get("/id/{id:int isPair()}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		fmt.Println(id)
	})

	app.Run(iris.Addr(":8080"))
}
