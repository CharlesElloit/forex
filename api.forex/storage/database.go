package storage

import (
	"log"
	"os"

	"api.forex.com/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectToDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		content, err := os.ReadFile(os.Getenv("DATABASE_URL_FILE"))
		if err != nil {
			log.Fatal(err)
		}
		dsn = string(content)
	}
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatalf("⛔ Unable to connect to database: %v\n", dbErr)
	} else {
		log.Println("Database connected... 🥇")
	}

	DB = db
	return db
}

func performMigrations(db *gorm.DB) {
	db.AutoMigrate(
		// add more models here
		&models.Fx_GL{},
		&models.Fx_User{},
		&models.Fx_Rate{},
		&models.Fx_Branch{},
		&models.Fx_GL_Type{},
		&models.Fx_GL_Total{},
		&models.Fx_Currency{},
		&models.Fx_Transaction{},
		&models.Fx_Charges_Type{},
		&models.Fx_Currency_Deno{},
	)
}

func InitializeDB() *gorm.DB {
	db := connectToDB()
	performMigrations(db)
	return db
}
