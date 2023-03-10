package db

import (
	"log"

	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// usar un .env
var DSN = "user=postgres password=4q2mQqUxfqnpgySz host=db.pzokydsbdlgmlwdtuirt.supabase.co port=5432 dbname=postgres"
var DB *gorm.DB

func DBConnection() {

	sqlDB, dbConnError := sql.Open("pgx", DSN)
	if dbConnError != nil {
		log.Fatal(dbConnError)
	}

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
		
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB successfully connected")
	}
}
