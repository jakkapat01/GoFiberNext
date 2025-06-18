package config

import (
	"fmt"
	"log"
	"os"

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

	log.Println("Database connected successfully")
	if shouldRunMigrations() {
		runMigration(db)
		// Seed admin user หลังจาก migration เสร็จ
		if err := SeedAdminUser(db, config); err != nil {
			log.Printf("Admin seeding failed: %v", err)
		}
	} else {
		autoMigrate := os.Getenv("AUTO_MIGRATE")
		appEnv := os.Getenv("APP_ENV")

		if autoMigrate == "false" {
			log.Println("Skipping database migrations as (AUTO_MIGRATE=false)")
		} else if appEnv == "production" && autoMigrate != "true" {
			log.Println("Skipping database migrations in (production environment, set AUTO_MIGRATE=true tp enable)")
		} else {
			log.Println("Skipping database migrations (set AUTO_MIGRATE=true to enable)")
		}
		// ลองสร้าง admin user แม้ว่าจะไม่ได้ migrate (กรณีที่ตารางมีอยู่แล้ว)
		if err := SeedAdminUser(db, config); err != nil {
			log.Printf("Admin seeding failed: %v", err)
		}
	}

	return db
}

// function ตรวจสอบว่าควร migrate database หรือไม่
func shouldRunMigrations() bool {
	if os.Getenv("AUTO_MIGRATE") == "false" {
		return false
	}
	if os.Getenv("AUTO_MIGRATE") == "true" {
		return true
	}
	// development environment ควร migrate อัตโนมัติ
	if os.Getenv("APP_ENV") == "development" {
		return true
	}
	// production environment ไม่ควร migrate อัตโนมัติ
	return false
}

func runMigration(db *gorm.DB) {
	log.Println("Starting database migrations...")
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Fail to migrate database: %v", err)
	}
	log.Println("Database migrations completed successfully")
}

// ฟังก์ชันสำหรับ migrate แบบ manual
func RunMigrationManual(config *Config) error {
	db := SetupDatabase(config)
	log.Println("Running manual database migration...")
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	log.Println("Manual database migration completed successfully")
	return nil
}
