package controller

import (
	"TanAgah/internal/config"
	"TanAgah/internal/entity"
	"TanAgah/internal/model"
	"TanAgah/internal/service"
	"TanAgah/internal/stringResource"
	"TanAgah/internal/utils"
	"github.com/gin-gonic/gin"
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
		utils.SendError400Response(ctx, err.Error())
		return
	}

	user := entity.User{
		Name:     registerRq.Name,
		Email:    registerRq.Email,
		Password: registerRq.Password,
		Role:     config.RoleUser,
	}
	if err := c.userService.CreateUser(&user); err != nil {
		utils.SendDataError500(ctx, err.Error())
		return
	}

	utils.SendSuccessResponse(ctx, model.MainRp{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		JwtToken: user.JwtToken,
	}, nil)
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	var LoginRq model.LoginRq
	if err := ctx.ShouldBindJSON(&LoginRq); err != nil {
		utils.SendError400Response(ctx, err.Error())
		return
	}

	user, err := c.userService.LoginUser(LoginRq.Email, LoginRq.Password)
	if err != nil {
		utils.SendDataError500(ctx, err.Error())
		return
	}

	utils.SendSuccessResponse(ctx, model.MainRp{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		JwtToken: user.JwtToken,
	}, nil)
	
}

func (c *UserController) DeleteUser(ctx *gin.Context) {

	id := ctx.Param("id")
	uintId, _ := strconv.ParseUint(id, 10, 32)

	var DeleteRq model.DeleteRq
	if err := ctx.ShouldBindJSON(&DeleteRq); err != nil {
		utils.SendError400Response(ctx, err.Error())
		return
	}

	user, err := c.userService.GetUserByUsername(DeleteRq.Email)
	if err != nil || user.Password != DeleteRq.Password || uint(uintId) != user.ID {
		utils.SendDataError403(ctx, stringResource.GetStrings().UnknownError(ctx))
		return
	}

	if err := c.userService.DeleteUser(uint(uintId)); err != nil {
		utils.SendDataError500(ctx, err.Error())
		return
	}
	utils.SendSuccessResponse(ctx, stringResource.GetStrings().UserDeleteSuccess(ctx), nil)
}
