package controller

import (
	"blog_web/db/service"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EssayController struct {
	essayService *service.EssayService
}

func NewEssayRouter() *EssayController {
	return &EssayController{
		essayService: service.NewEssayService(),
	}
}

func (e *EssayController) EssayList(ctx *gin.Context) {
	essays := e.essayService.GetAll()
	if essays == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}
	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, essays))
}