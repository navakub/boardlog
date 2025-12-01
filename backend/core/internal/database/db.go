package database

import (
	"fmt"
	"log"

	"github.com/navakub/boardlog/backend/core/internal/config"
	"github.com/navakub/boardlog/backend/core/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)
	if dsn == "" {
		log.Fatal("Database connection string is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	// Assign global DB
	DB = db
	log.Println("Database connected!")

	runMigrations()
}

func runMigrations() {
	err := DB.AutoMigrate(
		&model.User{},
		&model.PlayLog{},
	)

	if err != nil {
		log.Fatal("Migration error: ", err)
	}

	log.Println("Database migrated!")
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database not initialized. Call ConnectDatabase() first.")
	}
	return DB
}
