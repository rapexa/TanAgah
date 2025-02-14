package controller

import (
	"TanAgah/internal/config"
	"TanAgah/internal/entity"
	"TanAgah/internal/model"
	"TanAgah/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
	var registerRq model.RegisterRq
	if err := ctx.ShouldBindJSON(&registerRq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entity.User{
		Name:     registerRq.Name,
		Email:    registerRq.Email,
		Password: registerRq.Password,
		Role:     config.RoleUser,
	}
	if err := c.userService.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	var LoginRq model.LoginRq
	if err := ctx.ShouldBindJSON(&LoginRq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userService.LoginUser(LoginRq.Email, LoginRq.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.userService.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var input entity.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the existing user by ID
	existingUser, err := c.userService.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update the existing user with the new data
	existingUser.Name = input.Name
	existingUser.Email = input.Email
	existingUser.Password = input.Password

	// Save the updated user
	if err := c.userService.UpdateUser(existingUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": existingUser})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	uintId, _ := strconv.ParseUint(id, 10, 32)
	if err := c.userService.DeleteUser(uint(uintId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "User deleted successfully"})
}
