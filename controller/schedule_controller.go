package controller

import (
	"gin-ayo/dto"
	"gin-ayo/pkg/utils"
	"gin-ayo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ScheduleController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Detail(ctx *gin.Context)
}

type scheduleController struct {
	scheduleService service.ScheduleService
}

func (a scheduleController) Detail(ctx *gin.Context) {
	var DTODetailSchedule dto.DetailSchedule
	errReg := ctx.ShouldBindUri(&DTODetailSchedule)
	if errReg != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errReg.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response, err := a.scheduleService.Detail(ctx, DTODetailSchedule)

	if err != nil {
		res := utils.BuildErrorResponse("Failed To Get Detail Schedule", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(utils.CodeError(err.Error()), res)
		return
	} else {
		res := utils.BuildResponse(true, "OK", response)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a scheduleController) Delete(ctx *gin.Context) {
	var DTODeleteSchedule dto.DeleteSchedule
	errReg := ctx.ShouldBindUri(&DTODeleteSchedule)
	if errReg != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errReg.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response, err := a.scheduleService.Delete(ctx, DTODeleteSchedule)

	if err != nil {
		res := utils.BuildErrorResponse("Failed To Update Schedule", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(utils.CodeError(err.Error()), res)
		return
	} else {
		res := utils.BuildResponse(true, "OK", response)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a scheduleController) Update(ctx *gin.Context) {
	var DTOUpdateSchedule dto.UpdateSchedule
	errReg := ctx.ShouldBind(&DTOUpdateSchedule)
	if errReg != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errReg.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response, err := a.scheduleService.Update(ctx, DTOUpdateSchedule)

	if err != nil {
		res := utils.BuildErrorResponse("Failed To Update Schedule", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(utils.CodeError(err.Error()), res)
		return
	} else {
		res := utils.BuildResponse(true, "OK", response)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a scheduleController) Create(ctx *gin.Context) {
	var DTOCreateSchedule dto.CreateSchedule
	errReg := ctx.ShouldBind(&DTOCreateSchedule)
	if errReg != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errReg.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response, err := a.scheduleService.Create(ctx, DTOCreateSchedule)

	if err != nil {
		res := utils.BuildErrorResponse("Failed To Create Schedule", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(utils.CodeError(err.Error()), res)
		return
	} else {
		res := utils.BuildResponse(true, "OK", response)
		ctx.JSON(http.StatusOK, res)
	}
}

func NewScheduleController(scheduleServ service.ScheduleService) ScheduleController {
	return &scheduleController{scheduleService: scheduleServ}
}
