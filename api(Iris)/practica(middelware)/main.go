package main

import "github.com/kataras/iris"

/*
* Nota:{
* 	-> Se puede usar un middelware global y otro para uso especifico.
* 	-> Tambien se puede combinar ambos middelware
*   -> Podemos usar middelware al final de la funcion de una ruta especifica
*		y esta se puede tener globalmente o independientemente, se ejecuta con ctx.Done()
*		y para que se ejecute, la funcion tiene que tener el metodo Nect() para poder llamarla
*}
 */

func main() {
	app := iris.New()

	// tambien se puede usar de esta manera, pero se auto ejeutaria cuando vayamos a una ruta
	app.Use(middelG) // para usar la funcion que creamos -> esta opcion se usa para usar el middelware global en todas las funciones
	//app.Done(avant2) // se ejecuta luego del llamado a la ruta

	app.Get("/", middel, func(ctx iris.Context) {
		ctx.HTML("<h3>page initial</h3>")
		//ctx.Next() // para que se ejecute en Done se necesita colocar Next()
	})

	app.Get("/test", func(ctx iris.Context) {
		ctx.HTML("<h3>page 2</h3>")
	})

	app.Get("/testt", middel, func(ctx iris.Context) {
		ctx.HTML("<h3>page 3</h3>")
		ctx.Next() // si no se coloca next, el segundo middelware

	}, middel2)

	app.Run(iris.Addr(":8000"))
}

func middelG(ctx iris.Context) { // para evaluar parametros, middelwares
	ctx.HTML("<h1> Middelware global! </h1>")
	ctx.Next() // para que pueda seguir con la ejecucion de la funcion de la ruta
}

func middel(ctx iris.Context) { // para evaluar parametros, middelwares
	ctx.HTML("<h1> Middelware! </h1>")
	ctx.Next()
}

func middel2(ctx iris.Context) { // para evaluar parametros, middelwares
	ctx.HTML("<h1> fin del sitio! </h1>")
}
