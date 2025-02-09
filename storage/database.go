/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package storage

import (
	"log"
	"os"

	"api.forex.com/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {
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
		&models.GL{},
		&models.Role{},
		&models.User{},
		&models.Rate{},
		&models.Module{},
		&models.Branch{},
		&models.GlType{},
		&models.GlTotal{},
		&models.UserRole{},
		&models.Resource{},
		&models.Currency{},
		&models.Transaction{},
		&models.ChargesType{},
		&models.CurrencyDeno{},
		&models.PasswordPolicy{},
	)
}

func InitializeDB() *gorm.DB {
	db := ConnectToDB()
	performMigrations(db)
	AssignAdminAndUserRole(db)
	return db
}
