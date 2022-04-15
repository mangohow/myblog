package admin

import (
	"blog_web/db/service"
	"blog_web/model"
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

func (e *EssayController) EssayList(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")
	essays, err := e.essayService.GetLimited(pageNum, pageSize)
	if checkError(err, "Get essays error") {
		queryFailed(ctx)
		return
	}
	count, _ := e.essayService.GetCount()
	result := utils.ResponseResult(utils.QUERY_SUCCESS, essays)
	result["count"] = count
	ctx.JSON(http.StatusOK, result)
}

func (e *EssayController) AddEssay(ctx *gin.Context) {
	var essay model.Essay
	err := ctx.ShouldBind(&essay)
	if checkError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	essay.CreateTime = time.Now()
	err = e.essayService.AddEssay(&essay)
	if checkError(err, "Add essay error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}

func (e *EssayController) DeleteEssay(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	if (id <= 0) {
		deleteFailed(ctx)
		return
	}
	err := e.essayService.DeleteEssay(id)
	if checkError(err, "Delete essay error") {
		deleteFailed(ctx)
		return
	}

	deleteSuccess(ctx)
}

func (e *EssayController) UpdateEssay(ctx *gin.Context) {
	var essay model.Essay
	err := ctx.ShouldBind(&essay)
	if checkError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = e.essayService.UpdateEssay(&essay)
	if checkError(err, "Update essay error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}