package utilities

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func CreateError(statusCode int, title string, detail string, ctx iris.Context) {
	ctx.StopWithProblem(statusCode, iris.NewProblem().Title(title).Detail(detail))
}

func InternalServerError() {
	fmt.Println("Internal Server Error")
}
