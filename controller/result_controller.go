package controller

import (
	"gin-ayo/dto"
	"gin-ayo/pkg/utils"
	"gin-ayo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResultController interface {
	Create(ctx *gin.Context)
}

type resultController struct {
	resultService service.ResultService
}

func (a resultController) Create(ctx *gin.Context) {
	var DTOCreateResult dto.CreateResult
	errReg := ctx.ShouldBind(&DTOCreateResult)
	if errReg != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errReg.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response, err := a.resultService.Create(ctx, DTOCreateResult)

	if err != nil {
		res := utils.BuildErrorResponse("Failed To Create Result", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(utils.CodeError(err.Error()), res)
		return
	} else {
		res := utils.BuildResponse(true, "OK", response)
		ctx.JSON(http.StatusOK, res)
	}
}

func NewResultController(resultServ service.ResultService) ResultController {
	return &resultController{resultService: resultServ}
}
