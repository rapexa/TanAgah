package controller

import (
	"TanAgah/internal/entity"
	"TanAgah/internal/service"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockUserService is a mock implementation of the UserService interface
type MockUserService struct {
	CreateUserFunc func(user *entity.User) error
	GetUserFunc    func(id string) (*entity.User, error)
	UpdateUserFunc func(user *entity.User) error
	DeleteUserFunc func(id uint) error
}

func (m *MockUserService) CreateUser(user *entity.User) error {
	return m.CreateUserFunc(user)
}

func (m *MockUserService) GetUser(id string) (*entity.User, error) {
	return m.GetUserFunc(id)
}

func (m *MockUserService) UpdateUser(user *entity.User) error {
	return m.UpdateUserFunc(user)
}

func (m *MockUserService) DeleteUser(id uint) error {
	return m.DeleteUserFunc(id)
}

func setupRouter(userService service.UserService) *gin.Engine {
	router := gin.Default()
	userController := NewUserController(userService)

	api := router.Group("/api/v1")
	{
		api.POST("/users", userController.CreateUser)
		api.GET("/users/:id", userController.GetUser)
		api.PUT("/users/:id", userController.UpdateUser)
		api.DELETE("/users/:id", userController.DeleteUser)
	}

	return router
}

func TestCreateUser(t *testing.T) {
	// Mock UserService
	mockUserService := &MockUserService{
		CreateUserFunc: func(user *entity.User) error {
			user.ID = 1 // Simulate a created user with ID 1
			return nil
		},
	}

	// Setup router
	router := setupRouter(mockUserService)

	// Create a new user
	user := entity.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
	}
	userJSON, _ := json.Marshal(user)

	// Create a request
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]entity.User
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, uint(1), response["data"].ID)
	assert.Equal(t, "John Doe", response["data"].Name)
	assert.Equal(t, "john.doe@example.com", response["data"].Email)
}

func TestGetUser(t *testing.T) {
	// Mock UserService
	mockUserService := &MockUserService{
		GetUserFunc: func(id string) (*entity.User, error) {
			return &entity.User{
				ID:    1,
				Name:  "John Doe",
				Email: "john.doe@example.com",
			}, nil
		},
	}

	// Setup router
	router := setupRouter(mockUserService)

	// Create a request
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/1", nil)

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]entity.User
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, uint(1), response["data"].ID)
	assert.Equal(t, "John Doe", response["data"].Name)
	assert.Equal(t, "john.doe@example.com", response["data"].Email)
}

func TestUpdateUser(t *testing.T) {
	// Mock UserService
	mockUserService := &MockUserService{
		GetUserFunc: func(id string) (*entity.User, error) {
			return &entity.User{
				ID:    1,
				Name:  "John Doe",
				Email: "john.doe@example.com",
			}, nil
		},
		UpdateUserFunc: func(user *entity.User) error {
			return nil
		},
	}

	// Setup router
	router := setupRouter(mockUserService)

	// Update user data
	updatedUser := entity.User{
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
	}
	updatedUserJSON, _ := json.Marshal(updatedUser)

	// Create a request
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/1", bytes.NewBuffer(updatedUserJSON))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]entity.User
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Jane Doe", response["data"].Name)
	assert.Equal(t, "jane.doe@example.com", response["data"].Email)
}

func TestDeleteUser(t *testing.T) {
	// Mock UserService
	mockUserService := &MockUserService{
		DeleteUserFunc: func(id uint) error {
			return nil
		},
	}

	// Setup router
	router := setupRouter(mockUserService)

	// Create a request
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/users/1", nil)

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "User deleted successfully", response["data"])
}
