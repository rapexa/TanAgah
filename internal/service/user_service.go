package service

import (
	"TanAgah/internal/entity"
	"TanAgah/internal/repository"
	"TanAgah/internal/utils"
	"strconv"
)

type UserService interface {
	CreateUser(user *entity.User) error
	GetUser(id string) (*entity.User, error) // id is string
	GetUserByUsername(username string) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) error
	LoginUser(username, password string) (*entity.User, error)
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

func (s *userService) LoginUser(username, password string) (*entity.User, error) {

	user, err := s.userRepo.FindByUsername(username)
	if err != nil || !utils.VerifyPassword(user.Password, password) {
		return nil, err
	}
	token, err := utils.GenerateJWT(int(user.ID), username)
	user, err = s.userRepo.UpdateJwtTokenUser(username, token)

	return user, err
}

func (s *userService) GetUser(id string) (*entity.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return s.userRepo.FindByID(uint(userID))
}

func (s *userService) GetUserByUsername(username string) (*entity.User, error) {
	return s.userRepo.FindByUsername(username)
}

func (s *userService) UpdateUser(user *entity.User) error {
	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
