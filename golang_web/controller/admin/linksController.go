package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LinksController struct {
	linkService *service.LinkService
}

func NewLinksRouter() *LinksController {
	return &LinksController{
		linkService: service.NewLinkService(),
	}
}

func (l *LinksController) LinksList(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")

	links := l.linkService.GetLimitedLinks(pageNum, pageSize)
	categories := l.linkService.GetAllCategory()
	if links == nil || categories == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}
	count := l.linkService.GetLinkCount()
	result := utils.ResponseResult(utils.QUERY_SUCCESS, links)
	result["categories"] = categories
	result["count"] = count
	ctx.JSON(http.StatusOK, result)
}

func (l *LinksController) DeleteLink(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	err := l.linkService.DeleteLink(id)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_SUCCESS))
}

func (l *LinksController) UpdateLink(ctx *gin.Context) {
	var link model.Link
	err := ctx.ShouldBind(&link)
	if err != nil {
		utils.Logger().Warning("bind param error:%v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = l.linkService.UpdateLink(&link)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}

func (l *LinksController) AddLink(ctx *gin.Context) {
	var link model.Link
	err := ctx.ShouldBind(&link)
	if err != nil {
		utils.Logger().Warning("bind param error:%v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = l.linkService.AddLink(&link)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}

func (l *LinksController) Categories(ctx *gin.Context) {
	categories := l.linkService.GetAllCategory()
	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, categories))
}

func (l *LinksController) DeleteCategory(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	err := l.linkService.DeleteCategory(id)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_SUCCESS))
}

func (l *LinksController) UpdateCategory(ctx *gin.Context) {
	var category model.LinkCategory
	err := ctx.ShouldBind(&category)
	if err != nil {
		utils.Logger().Warning("bind param error:%v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = l.linkService.UpdateCategory(&category)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}

func (l *LinksController) AddCategory(ctx *gin.Context) {
	var category model.LinkCategory
	err := ctx.ShouldBind(&category)
	if err != nil {
		utils.Logger().Warning("bind param error:%v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = l.linkService.AddCategory(&category)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}