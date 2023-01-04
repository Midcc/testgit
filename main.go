package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	app.Get("/iris", func(ctx iris.Context) {

		ctx.WriteString("hello iris")
	})
	app.Listen(":8080")
}
