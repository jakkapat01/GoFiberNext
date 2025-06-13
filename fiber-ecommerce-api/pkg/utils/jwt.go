package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, role string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := &Claims{
		UserID: fmt.Sprintf("%d", userID),
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // Token issued at current time
			Issuer:    "fiber-ecommerce-api",                              // Issuer of the token
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ฟังก์ชัน ValidateJWT ใช้เพื่อตรวจสอบความถูกต้องของ JWT token
func ValidateJWT(tokenString string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	// ตรวจสอบว่ามีข้อผิดพลาดในการวิเคราะห์ token หรือไม่
	if err != nil {
		return nil, err
	}
	// ตรวจสอบว่า token มี claims ที่ถูกต้องและเป็น valid หรือไม่
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	// หาก token ไม่ valid หรือ claims ไม่ถูกต้อง ให้คืนค่า nil และ error
	return nil, jwt.ErrSignatureInvalid
}
