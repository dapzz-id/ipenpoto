package config

import (
	// "ipenpoto/database/models"
	"github.com/gofiber/fiber/v2"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database.\n", err)
	}

	// Ambil instance sql.DB dari GORM
	sqlDB, err := DB.DB()
	
	if err != nil {
		log.Fatal("Failed to get sql.DB instance.\n", err)
	}

	// Ambil nilai dari .env
	maxOpenConns := 50
	maxIdleConns := 10
	connMaxLifetime := 300


	if v := os.Getenv("DB_MAX_OPEN_CONNS"); v != "" {
		fmt.Sscanf(v, "%d", &maxOpenConns)
	}

	if v := os.Getenv("DB_MAX_IDLE_CONNS"); v != "" {
		fmt.Sscanf(v, "%d", &maxIdleConns)
	}

	if v := os.Getenv("DB_CONN_MAX_LIFETIME"); v != "" {
		fmt.Sscanf(v, "%d", &connMaxLifetime)
	}

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	log.Println("Database connection successfully opened")
}

func MigrateDB() {
	if fiber.IsChild() {
		return
	}

	// =================================
	// 1. AUTO MIGRATE TABLES (NO CONSTRAINTS)
	// =================================
	err := DB.AutoMigrate(
		//
	)

	if err != nil {
		log.Fatal("Failed to auto migrate database.\n", err)
	}

	log.Println("Database migrated successfully with physical relations")
}
