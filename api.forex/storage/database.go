package storage

import (
	"log"
	"os"

	"api.forex.com/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectToDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Panic("Fail to connection to the database.")
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
