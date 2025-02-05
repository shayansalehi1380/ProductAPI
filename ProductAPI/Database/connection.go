package Database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBHost     = "localhost"
	DBPort     = "5432"
	DBUser     = "postgres"
	DBPassword = "09011155"
	DBName     = "url_shortener"
)

var DB *gorm.DB

func ConnectDatabase() {
    // ایجاد رشته اتصال (DSN) برای PostgreSQL
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        DBHost, DBUser, DBPassword, DBName, DBPort)

    // اتصال به پایگاه داده PostgreSQL
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to the database!", err)
    }

    DB = db
    log.Println("Database connected...")
}