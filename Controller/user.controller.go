package controller

import (
	"github.com/gin-gonic/gin"
	helper "github.com/vandenbill/brand-commerce/Helper"
	web "github.com/vandenbill/brand-commerce/Model/Web"
	service "github.com/vandenbill/brand-commerce/Service"
	"net/http"
)

type UserController interface {
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type userControllerImpl struct {
	userService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &userControllerImpl{userService: userService}
}

func (u *userControllerImpl) Update(ctx *gin.Context) {
	sub := helper.GetSub(ctx)

	var userDto web.UserDto
	ctx.ShouldBind(&userDto)

	user, err := u.userService.Update(sub, userDto)

	if err != nil {
		res := helper.BuildErrorResponse("Failed to edit user", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	userRes := web.UserResponse{
		Id:   user.ID,
		Name: user.Name,
	}
	res := helper.BuildResponse("Succes update user", userRes)
	ctx.JSON(http.StatusCreated, res)
}

func (u *userControllerImpl) Delete(ctx *gin.Context) {
	sub := helper.GetSub(ctx)

	err := u.userService.Delete(sub)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse("Succes delete user", gin.H{})
	ctx.JSON(http.StatusOK, res)
}

func (u *userControllerImpl) FindAll(ctx *gin.Context) {
	users, err := u.userService.FindAll()
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all users data", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse("Succes get all users data", users)
	ctx.JSON(http.StatusOK, res)
}

func (u *userControllerImpl) FindByID(ctx *gin.Context) {
	sub := helper.GetSub(ctx)

	user, err := u.userService.FindByID(sub)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get user data", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse("Succes get user data", user)
	ctx.JSON(http.StatusOK, res)
}
