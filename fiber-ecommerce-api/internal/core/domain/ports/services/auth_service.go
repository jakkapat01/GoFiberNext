package services

import (
	"github.com/jakkapat01/fiber-ecommerce-api/internal/core/domain/entities"
)

type AuthService interface {
	Register(req entities.RegisterRequest) (*entities.User, error)
	AdminRegister(req entities.AdminRegisterRequest) (*entities.User, error)
	Login(req entities.LoginRequest) (*entities.LoginResponse, error)
	GetUserByID(id uint) (*entities.User, error)
	UpdateUser(user *entities.User) error
}
