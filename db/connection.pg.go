package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func DBConnection() {

	db_Host := os.Getenv("POST_DB_HOST")
	db_Port := os.Getenv("POST_DB_PORT")
	db_Name := os.Getenv("POST_DB_NAME")
	db_User := os.Getenv("POST_DB_USER")
	db_Password := os.Getenv("POST_DB_PASSWORD")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", db_User, db_Password, db_Host, db_Port, db_Name)

	sqlDB, dbConnError := sql.Open("pgx", dsn)
	if dbConnError != nil {
		log.Println(dbConnError)
	}
	log.Printf("%s DB connection established.\n", db_Name)

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println(err)
	}

	log.Println("DB successfully connected")
}
