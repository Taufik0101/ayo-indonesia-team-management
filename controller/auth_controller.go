package controller

import (
	"gin-ayo/dto"
	"gin-ayo/pkg/utils"
	"gin-ayo/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func (a authController) Login(ctx *gin.Context) {
	var DTOLogin dto.LoginUser
	errReg := ctx.ShouldBind(&DTOLogin)
	logrus.Println(errReg)
	if errReg != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errReg.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response, err := a.authService.Login(DTOLogin)

	if err != nil {
		res := utils.BuildErrorResponse("Login Failed", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(utils.CodeError(err.Error()), res)
		return
	} else {
		res := utils.BuildResponse(true, "OK", response)
		ctx.JSON(http.StatusOK, res)
	}
}

func NewAuthController(authServ service.AuthService) AuthController {
	return &authController{authService: authServ}
}
