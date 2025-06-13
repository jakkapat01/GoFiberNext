package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword ใช้สำหรับแฮชรหัสผ่านด้วย bcrypt
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CheckPasswordHash ใช้สำหรับตรวจสอบรหัสผ่านที่แฮชแล้วกับรหัสผ่านที่ผู้ใช้ป้อน
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
