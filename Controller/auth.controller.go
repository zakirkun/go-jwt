package controller

import (
	"github.com/gin-gonic/gin"
	helper "github.com/vandenbill/brand-commerce/Helper"
	web "github.com/vandenbill/brand-commerce/Model/Web"
	service "github.com/vandenbill/brand-commerce/Service"
	"net/http"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authControllerImpl struct {
	authService service.AuthService
	userService service.UserService
	jwtService  service.JwtService
}

func NewAuthControllerImpl(authService service.AuthService, userService service.UserService, jwtService service.JwtService) AuthController {
	return &authControllerImpl{authService: authService, userService: userService, jwtService: jwtService}
}

func (a authControllerImpl) Register(ctx *gin.Context) {
	var createUserDto web.UserDto
	err := ctx.ShouldBind(&createUserDto)

	if err != nil {
		res := helper.BuildErrorResponse("Failed procces request", "Can't bind dto", "Request can't be recorded", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	user, err := a.userService.Create(createUserDto)
	if err != nil {
		res := helper.BuildErrorResponse("Failed save data to database", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response := helper.BuildResponse("Register succes", web.UserResponse{
		Id:   user.ID,
		Name: user.Name,
	})
	ctx.JSON(http.StatusCreated, response)
}

// TODO should be in service
func (a authControllerImpl) Login(ctx *gin.Context) {
	var loginDto web.LoginDto
	err := ctx.ShouldBind(&loginDto)

	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = a.authService.VerifyCredential(loginDto)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to login", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, _ := a.userService.FindByEmail(loginDto.Email)

	token, err := a.jwtService.GenerateToken(user.ID)

	if err != nil {
		response := helper.BuildErrorResponse("Failed to sign token", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("acces_token", token, 86400, "", "", false, true)
	response := helper.BuildResponse("Login succes", helper.Obj{})
	ctx.JSON(http.StatusOK, response)
}
