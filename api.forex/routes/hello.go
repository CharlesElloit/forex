package routes

import "github.com/kataras/iris/v12"

func HelloWorld(ctx iris.Context) {
	ctx.Text("Hello World")
}
