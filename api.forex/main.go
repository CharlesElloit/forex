package main

import (
	"os"

	"api.forex.com/routes"
	"api.forex.com/storage"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	godotenv.Load()
	storage.InitializeDB()
	app := iris.Default()

	forex_rate := app.Party("/api/hello")
	{
		forex_rate.Get("/", routes.HelloWorld)
	}

	contacts := app.Party("/api/contacts")
	{
		contacts.Get("/all", routes.GetContacts)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	app.Listen(":" + port)
}
