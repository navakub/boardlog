package service

import (
	"github.com/navakub/boardlog/backend/core/internal/model"
	"github.com/navakub/boardlog/backend/core/internal/repository"
	"github.com/navakub/boardlog/backend/core/internal/utils"

	// "github.com/navakub/boardlog/backend/core/internal/middleware"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(user *model.User) error
	Login(email, password string) (*model.User, error)
	Logout(userID uint) error
	GetCurrentUser(userID uint) (*model.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(user *model.User) error {
	return s.userRepo.Create(user)
}

func (s *authService) Login(email, password string) (*model.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (s *authService) Logout(userID uint) error {
	// Implement logout logic if needed (e.g., token invalidation)
	return nil
}

func (s *authService) GetCurrentUser(userID uint) (*model.User, error) {
	return s.userRepo.GetByID(userID)
}
