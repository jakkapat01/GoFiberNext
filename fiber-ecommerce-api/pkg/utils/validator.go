package utils

import (
	"github.com/go-playground/validator/v10"
)

// สร้างตัวแปรสำหรับ validator
var validate = validator.New()

// สร้างฟังก์ชันสำหรับการตรวจสอบความถูกต้องของข้อมูล
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
