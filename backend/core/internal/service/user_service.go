package service

import (
	"github.com/navakub/boardlog/backend/core/internal/model"
	"github.com/navakub/boardlog/backend/core/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.userRepo.GetAll()
}

func (s *userService) GetUserByID(id uint) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	return s.userRepo.GetByEmail(email)
}

func (s *userService) CreateUser(user *model.User) error {
	return s.userRepo.Create(user)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
