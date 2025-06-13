package repositories

import (
	"github.com/jakkapat01/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/jakkapat01/fiber-ecommerce-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

// UserRepositoryImpl เป็นการใช้งาน GORM เพื่อจัดการกับฐานข้อมูลสำหรับ User
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

// Create สร้างผู้ใช้ใหม่ในฐานข้อมูล
func (r *UserRepositoryImpl) Create(user *entities.User) error {
	userModel := &models.User{}
	userModel.FromEntity(user)
	if err := r.db.Create(userModel).Error; err != nil {
		return err
	}
	*user = *userModel.ToEntity()
	return nil
}

// GetByEmail ค้นหาผู้ใช้ตามอีเมล
func (r *UserRepositoryImpl) GetByEmail(email string) (*entities.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
}

// GetByID ค้นหาผู้ใช้ตาม ID
func (r *UserRepositoryImpl) GetByID(id uint) (*entities.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
}

// Update อัปเดตข้อมูลผู้ใช้
func (r *UserRepositoryImpl) Update(user *entities.User) error {
	userModel := &models.User{}
	userModel.FromEntity(user)
	return r.db.Save(userModel).Error
}

// Delete ลบผู้ใช้ตาม ID
func (r *UserRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// GetAll ดึงข้อมูลผู้ใช้ทั้งหมด
func (r *UserRepositoryImpl) GetAll() ([]*entities.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	var result []*entities.User
	for _, user := range users {
		result = append(result, user.ToEntity())
	}
	return result, nil
}
