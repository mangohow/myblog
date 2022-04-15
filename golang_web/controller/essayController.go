package controller

import (
	"blog_web/db/service"
	"github.com/gin-gonic/gin"
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
	essays, err := e.essayService.GetAll()
	if checkError(err, "Get Essay List") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, essays)
}