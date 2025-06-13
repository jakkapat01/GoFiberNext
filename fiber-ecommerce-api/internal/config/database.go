package config

import (
	"fmt"
	"log"

	"github.com/jakkapat01/fiber-ecommerce-api/internal/adapters/persistence/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort, config.DBSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}
	log.Println("Database connected and migrated successfully")

	return db
}
