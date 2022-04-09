package router

import (
	"blog_web/db/service"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EssayRouter struct {
	essayService *service.EssayService
}

func NewEssayRouter() *EssayRouter {
	return &EssayRouter{
		essayService: service.NewEssayService(),
	}
}

func (e *EssayRouter) EssayList(ctx *gin.Context) {
	essays := e.essayService.GetAll()
	if essays == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}
	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, essays))
}