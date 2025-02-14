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

func (c *UserController) DeleteUser(ctx *gin.Context) {

	id := ctx.Param("id")
	uintId, _ := strconv.ParseUint(id, 10, 32)

	var DeleteRq model.DeleteRq
	if err := ctx.ShouldBindJSON(&DeleteRq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userService.GetUserByUsername(DeleteRq.Email)
	if err != nil || user.Password != DeleteRq.Password {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	if err := c.userService.DeleteUser(uint(uintId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "User deleted successfully"})
}
