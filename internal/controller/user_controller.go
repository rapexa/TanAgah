package controller

import (
	"TanAgah/internal/entity"
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

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.userService.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	uintId, _ := strconv.ParseUint(id, 10, 32)
	user, err := c.userService.GetUser(uint(uintId))
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
	uintId, _ := strconv.ParseUint(id, 10, 32)
	existingUser, err := c.userService.GetUser(uint(uintId))
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
