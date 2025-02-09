package main

import (
	forexCommon "api.forex.com/forex.common"
	forexRate "api.forex.com/forex.rates/api"
	forexReport "api.forex.com/forex.report/rate"
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

	forex_rate := app.Party("/api/exchange/rate")
	{
		forex_rate.Post("/add", forexRate.AddNewExchangeRate)
		forex_rate.Get("/report/rates", forexReport.GetRate)
	}

	var (
		port = forexCommon.LoadEnv("PORT", ":3000")
	)

	app.Listen(":" + port)
}
