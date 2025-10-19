package controller

import (
	"gin-ayo/dto"
	"gin-ayo/pkg/utils"
	"gin-ayo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PlayerController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type playerController struct {
	playerService service.PlayerService
}

func (a playerController) Delete(ctx *gin.Context) {
	var DTODeletePlayer dto.DeletePlayer
	errReg := ctx.ShouldBindUri(&DTODeletePlayer)
	if errReg != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errReg.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response, err := a.playerService.Delete(ctx, DTODeletePlayer)

	if err != nil {
		res := utils.BuildErrorResponse("Failed To Update Player", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(utils.CodeError(err.Error()), res)
		return
	} else {
		res := utils.BuildResponse(true, "OK", response)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a playerController) Update(ctx *gin.Context) {
	var DTOUpdatePlayer dto.UpdatePlayer
	errReg := ctx.ShouldBind(&DTOUpdatePlayer)
	if errReg != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errReg.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response, err := a.playerService.Update(ctx, DTOUpdatePlayer)

	if err != nil {
		res := utils.BuildErrorResponse("Failed To Update Player", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(utils.CodeError(err.Error()), res)
		return
	} else {
		res := utils.BuildResponse(true, "OK", response)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a playerController) Create(ctx *gin.Context) {
	var DTOCreatePlayer dto.CreatePlayer
	errReg := ctx.ShouldBind(&DTOCreatePlayer)
	if errReg != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errReg.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response, err := a.playerService.Create(ctx, DTOCreatePlayer)

	if err != nil {
		res := utils.BuildErrorResponse("Failed To Create Player", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(utils.CodeError(err.Error()), res)
		return
	} else {
		res := utils.BuildResponse(true, "OK", response)
		ctx.JSON(http.StatusOK, res)
	}
}

func NewPlayerController(playerServ service.PlayerService) PlayerController {
	return &playerController{playerService: playerServ}
}
