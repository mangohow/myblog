package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type EssayController struct {
	essayService *service.EssayService
}

func NewEssayRouter() *EssayController {
	return &EssayController{
		essayService: service.NewEssayService(),
	}
}

func (e *EssayController) EssayList(ctx *gin.Context) *response.Response  {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")
	essays, err := e.essayService.GetLimited(pageNum, pageSize)
	if response.CheckError(err, "Get essays error") {
		return response.ResponseQueryFailed()
	}
	count, _ := e.essayService.GetCount()

	return response.ResponseQuerySuccess(essays, count)
}

func (e *EssayController) AddEssay(ctx *gin.Context) *response.Response {
	var essay model.Essay
	err := ctx.ShouldBind(&essay)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	essay.CreateTime = time.Now()
	err = e.essayService.AddEssay(&essay)
	if response.CheckError(err, "Add essay error") {
		return response.ResponseOperateFailed()
	}

	return response.ResponseOperateSuccess()
}

func (e *EssayController) DeleteEssay(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	if (id <= 0) {
		return response.ResponseDeleteFailed()
	}
	err := e.essayService.DeleteEssay(id)
	if response.CheckError(err, "Delete essay error") {
		return response.ResponseDeleteFailed()
	}

	return response.ResponseDeleteSuccess()
}

func (e *EssayController) UpdateEssay(ctx *gin.Context) *response.Response {
	var essay model.Essay
	err := ctx.ShouldBind(&essay)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	err = e.essayService.UpdateEssay(&essay)
	if response.CheckError(err, "Update essay error") {
		return response.ResponseOperateFailed()
	}

	return response.ResponseOperateSuccess()
}