package repository

import (
	"TanAgah/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByID(id uint) (*entity.User, error) // id is uint
	Update(user *entity.User) error
	Delete(id uint) error
	FindByUsername(username string) (*entity.User, error)
	UpdateJwtTokenUser(username string, jwtToken string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateJwtTokenUser(username string, jwtToken string) (*entity.User, error) {
	var user entity.User
	// Find the user by email (username)
	err := r.db.Where("email = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	// Assuming the User struct has a JwtToken field
	user.JwtToken = jwtToken

	// Save the updated user object
	err = r.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}
