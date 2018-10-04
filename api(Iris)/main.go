package main

import (
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

	app.Get("/param", func() {

	})

	app.Run(iris.Addr(":8080"))
}
