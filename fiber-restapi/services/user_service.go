package services

import (
	"fiber-restapi/models"
	"fmt"
	"time"
)

type UserService struct {
	users  []models.User // Simulating a database with a slice
	nextID int           // id ถัดไปสำหรับผู้ใช้ใหม่

}

// สร้าง function ใหม่สำหรับ UserService
func NewUserService() *UserService {
	return &UserService{
		users:  []models.User{},
		nextID: 1, // เริ่มต้น ID ถัดไปที่ 1

	}
}

// สร้าง function Register สำหรับลงทะเบียนผู้ใช้ใหม่
func (s *UserService) Register(user models.User) (*models.User, error) {
	// ตรวจสอบว่าผู้ใช้มีอยู่แล้วหรือไม่
	for _, u := range s.users {
		if u.Username == user.Username {
			return nil, fmt.Errorf("email %s already exists", user.Username)
		}
		if u.Email == user.Email {
			return nil, fmt.Errorf("email %s already exists", user.Email)
		}
	}
	user.ID = s.nextID
	s.nextID++                             // เพิ่ม ID ถัดไป
	now := time.Now().Format(time.RFC3339) // ใช้เวลาในรูป ISO 8601 เช่น "2023-10-01T12:00:00Z"
	user.CreatedAt = now
	user.UpdatedAt = now
	// กำหนดค่าเริ่มต้นสำหรับ role เป็น "user"
	if user.Role == "" {
		user.Role = "user"
	}
	// เพิ่มผู้ใช้ใหม่ลงใน slice
	s.users = append(s.users, user)
	return &user, nil

}

// สร้าง function สำหรับ login ผู้ใช้
func (s *UserService) Login(username, password string) (*models.User, string, error) {
	for _, u := range s.users {
		if u.Username == username && u.Password == password {
			// สร้าง token แบบง่ายๆ (ในงานจริงควรใช้ JWT)
			token := fmt.Sprintf("token_%s_%d", u.Username, time.Now().Unix())
			return &u, token, nil
		}
	}
	return nil, "", fmt.Errorf("invalid username or password")
}
