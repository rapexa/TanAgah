package service

import (
	"TanAgah/internal/entity"
	"TanAgah/internal/repository"
)

type UserService interface {
	CreateUser(user *entity.User) error
	GetUser(id uint) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) CreateUser(user *entity.User) error {
	return s.userRepo.Create(user)
}

func (s *userService) GetUser(id uint) (*entity.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) UpdateUser(user *entity.User) error {
	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
