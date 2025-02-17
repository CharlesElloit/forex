package main

import (
	forexAuth "api.forex.com/forex.auth/api"
	forexCommon "api.forex.com/forex.common"
	forexRate "api.forex.com/forex.rates/api"
	forexReport "api.forex.com/forex.report/api/rate"
	"api.forex.com/storage"
	"github.com/go-playground/validator"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kataras/iris/v12"
)

func main() {
	storage.InitializeDB()
	app := iris.Default()
	app.Validator = validator.New()

	healthCheck := app.Party("/ping")
	{
		healthCheck.Get("/", func(ctx iris.Context) {
			ctx.JSON(iris.StatusOK)
		})
	}

	auth := app.Party("/api/user/auth")
	{
		auth.Post("/login", forexAuth.Login)
	}

	forex_rate := app.Party("/api/exchange/rate")
	{
		forex_rate.Post("/add", forexRate.Add)
		forex_rate.Get("/report/rates", forexReport.GetRate)
	}

	var (
		port = forexCommon.LoadEnv("PORT", ":3000")
	)

	app.Listen(":" + port)
}
