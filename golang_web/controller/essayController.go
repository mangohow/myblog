package controller

import (
	"blog_web/db/service"
	"blog_web/response"
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

func (e *EssayController) EssayList(ctx *gin.Context) *response.Response {
	essays, err := e.essayService.GetAll()
	if response.CheckError(err, "Get Essay List") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(essays)
}