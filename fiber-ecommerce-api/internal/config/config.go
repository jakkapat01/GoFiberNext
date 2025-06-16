package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv       string
	AppPort      string
	AppURL       string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSSLMode    string
	JWTSecret    string
	JWTExpiresIn string
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func LoadConfig() (*Config, error) {

	// โหลดไฟล์ .env ถ้ามี
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	config := &Config{
		// ค่าที่ปลอดภัยจากการตั้งค่าเริ่มต้น
		AppEnv:       getEnv("APP_ENV", "development"),
		AppPort:      getEnv("APP_PORT", "3000"),
		AppURL:       getEnv("APP_URL", "http://localhost:3000"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBSSLMode:    getEnv("DB_SSL", "disable"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),
		// ค่าที่ต้องการให้ผู้ใช้กำหนดในไฟล์ .env
		DBPassword: getEnv("DB_PASS", ""),
		DBName:     getEnv("DB_NAME", ""),
		JWTSecret:  getEnv("JWT_SECRET", ""),
	}
	if err := validateConfig(config); err != nil {
		return nil, err
	}
	return config, nil
}

// ฟังก์ชันตรวจสอบว่าค่าที่จำเป็นถูกตั้งค่าในไฟล์ .env หรือไม่
func validateConfig(config *Config) error {
	// ตรวจสอบค่าที่จำเป็นสำหรับ production
	if config.AppEnv == "production" {
		if config.DBPassword == "" {
			return fmt.Errorf("DB_PASS is required in production environment")
		}
		if config.JWTSecret == "" {
			return fmt.Errorf("JWT_SECRET is required in production environment")
		}
		if len(config.JWTSecret) < 32 {
			return fmt.Errorf("JWT_SECRET must be at least 32 characters long for production security")
		}
		if config.DBSSLMode == "" {
			log.Println("Warning: SSL is disabled for database connection in production")
		}
	}
	if config.DBName == "" {
		return fmt.Errorf("DB_NAME is required")
	}
	return nil

}
