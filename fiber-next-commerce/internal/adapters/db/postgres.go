package db

import (
	"fmt"
	"log"

	"github.com/jakkapat01/fiber-next-commerce/config"
	"github.com/jakkapat01/fiber-next-commerce/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connectDB เชื่อมต่อกับฐานข้อมูล PostgreSQL
func Connect(config *config.Config) *gorm.DB {
	// สร้าง connection string สำหรับ PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s  sslmode=%s TimeZone=Asia/Bangkok", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort, config.DBSSL)
	// เชื่อมต่อกับฐานข้อมูล
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("ไม่สารถเชื่อมต่อฐานข้อมูลได้: %v", err)
	}

	// สร้างฐานข้อมูลจากโมเดล
	migrateDatabase(db)

	log.Println("เชื่อมต่อฐานข้อมูลสำเร็จ")

	return db
}

// MigrateDB ทำการ migrate schema ของฐานข้อมูล
func migrateDatabase(db *gorm.DB) {
	// AutoMigrate จะทำการสร้างตารางและคอลัมน์ที่จำเป็นตามโครงสร้างของโมเดล
	err := db.AutoMigrate(
		&domain.Role{},
		&domain.Permission{},
		&domain.User{},
		&domain.Category{},
		&domain.Product{},
		&domain.ProductImage{},
		&domain.Cart{},
		&domain.CartItem{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Transaction{},
	)
	if err != nil {
		log.Fatalf("ไม่สามารถสร้างตารางในฐานข้อมูลได้: %v", err)
	}
	log.Println("สร้างตารางในฐานข้อมูลสำเร็จ")
}
