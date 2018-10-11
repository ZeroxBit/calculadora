package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var sessionConfig = sessions.Config{Cookie: "session"} // se crear el nombre del la cookei
var sess = sessions.New(sessionConfig)

func main() {
	app := iris.New()

	app.Get("/login", func(ctx iris.Context) {
		sess.Start(ctx).Set("Authorization", true) // crear una cookie
		ctx.HTML("<h1>Login !!!</h1>")
	})

	app.Get("/logout", func(ctx iris.Context) {
		sess.Start(ctx).Set("Authorization", false) // borra una cookie
		ctx.HTML("<h1>Adios :(</h1>")
	})

	app.Get("/connect", func(ctx iris.Context) {
		author, _ := sess.Start(ctx).GetBoolean("Authorization")
		if !author { // redirecciona si tiene token o no
			ctx.Redirect("/disconnect")
		}

		ctx.HTML("<h1>connect :)</h1>")

	})

	app.Get("/disconnect", func(ctx iris.Context) {
		ctx.HTML("<h1>disConnect :(</h1>")
	})

	app.Get("/session", func(ctx iris.Context) {
		author, _ := sess.Start(ctx).GetBoolean("Authorization")
		if !author { // redirecciona si tiene token o no
			ctx.Redirect("/disconnect")
		} else {
			ctx.Redirect("/connect")
		}

	})

	app.Run(iris.Addr(":8000"))
}
