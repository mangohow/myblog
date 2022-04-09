package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")
	essays := e.essayService.GetLimited(pageNum, pageSize)
	count := e.essayService.GetCount()
	result := utils.ResponseResult(utils.QUERY_SUCCESS, essays)
	result["count"] = count
	ctx.JSON(http.StatusOK, result)
}

func (e *EssayRouter) AddEssay(ctx *gin.Context) {
	var essay model.Essay
	err := ctx.ShouldBind(&essay)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	essay.CreateTime = time.Now()
	err = e.essayService.AddEssay(&essay)
	if err != nil {
		utils.Logger().Warning("add essay error:%v", err)
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}

func (e *EssayRouter) DeleteEssay(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	if (id <= 0) {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_FAILED))
		return
	}
	err := e.essayService.DeleteEssay(id)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_SUCCESS))
}

func (e *EssayRouter) UpdateEssay(ctx *gin.Context) {
	var essay model.Essay
	err := ctx.ShouldBind(&essay)
	if err != nil {
		utils.Logger().Warning("bind essay error:%v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = e.essayService.UpdateEssay(&essay)
	if err != nil {
		utils.Logger().Warning("update essay error:%v", err)
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}