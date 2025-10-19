package controller

import (
	"gin-ayo/dto"
	"gin-ayo/pkg/utils"
	"gin-ayo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TeamController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type teamController struct {
	teamService service.TeamService
}

func (a teamController) Delete(ctx *gin.Context) {
	var DTODeleteTeam dto.DeleteTeam
	errReg := ctx.ShouldBindUri(&DTODeleteTeam)
	if errReg != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errReg.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response, err := a.teamService.Delete(ctx, DTODeleteTeam)

	if err != nil {
		res := utils.BuildErrorResponse("Failed To Update Team", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(utils.CodeError(err.Error()), res)
		return
	} else {
		res := utils.BuildResponse(true, "OK", response)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a teamController) Update(ctx *gin.Context) {
	var DTOUpdateTeam dto.UpdateTeam
	errReg := ctx.ShouldBind(&DTOUpdateTeam)
	if errReg != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errReg.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response, err := a.teamService.Update(ctx, DTOUpdateTeam)

	if err != nil {
		res := utils.BuildErrorResponse("Failed To Update Team", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(utils.CodeError(err.Error()), res)
		return
	} else {
		res := utils.BuildResponse(true, "OK", response)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a teamController) Create(ctx *gin.Context) {
	var DTOCreateTeam dto.CreateTeam
	errReg := ctx.ShouldBind(&DTOCreateTeam)
	if errReg != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errReg.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response, err := a.teamService.Create(ctx, DTOCreateTeam)

	if err != nil {
		res := utils.BuildErrorResponse("Failed To Create Team", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(utils.CodeError(err.Error()), res)
		return
	} else {
		res := utils.BuildResponse(true, "OK", response)
		ctx.JSON(http.StatusOK, res)
	}
}

func NewTeamController(teamServ service.TeamService) TeamController {
	return &teamController{teamService: teamServ}
}
