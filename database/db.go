package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Database(envpath string) *gorm.DB {

	if err := godotenv.Load(envpath); err != nil {
		// In production/K8s env vars are injected directly — .env file is optional.
		log.Printf("Warning: .env file not found at %s, using system environment variables", envpath)
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Database credentials are not fully set in the environment variables")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Retry loop — in K8s the API pod can start before Postgres is ready.
	// Retry every 3 seconds for up to 30 seconds before giving up.
	var db *gorm.DB
	var err error
	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			sqlDB, pingErr := db.DB()
			if pingErr == nil {
				if pingErr = sqlDB.Ping(); pingErr == nil {
					break
				}
			}
			err = pingErr
		}
		log.Printf("Waiting for database... attempt %d/10 (%v)", i, err)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Fatalf("Failed to connect to database after 10 attempts: %v", err)
	}

	log.Println("Database connected successfully")

	return db
}

func Migrations(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB for migrations: %v", err)
	}
	if err := RunMigrations(sqlDB); err != nil {
		log.Fatal(err)
	}
}

// SQLDb returns the underlying *sql.DB from a *gorm.DB (used for migrations).
func SQLDb(db *gorm.DB) *sql.DB {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}
	return sqlDB
}
