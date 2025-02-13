package service

import (
	"TanAgah/internal/entity"
	"TanAgah/internal/repository"
	"strconv"
)

type UserService interface {
	CreateUser(user *entity.User) error
	GetUser(id string) (*entity.User, error) // id is string
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

func (s *userService) GetUser(id string) (*entity.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return s.userRepo.FindByID(uint(userID))
}

func (s *userService) UpdateUser(user *entity.User) error {
	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
