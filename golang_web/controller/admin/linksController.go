package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
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

func (l *LinksController) LinksList(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")

	links, err := l.linkService.GetLimitedLinks(pageNum, pageSize)
	if response.CheckError(err, "Get links error") {
		return response.ResponseQueryFailed()
	}
	categories, err := l.linkService.GetAllCategory()
	if response.CheckError(err, "Get categories error") {
		return response.ResponseQueryFailed()
	}

	count, _ := l.linkService.GetLinkCount()

	return response.ResponseQuerySuccess(links, categories, count)
}

func (l *LinksController) DeleteLink(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	err := l.linkService.DeleteLink(id)
	if response.CheckError(err, "Delete link error") {
		return response.ResponseDeleteFailed()
	}

	return response.ResponseDeleteSuccess()
}

func (l *LinksController) UpdateLink(ctx *gin.Context) *response.Response {
	var link model.Link
	err := ctx.ShouldBind(&link)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	err = l.linkService.UpdateLink(&link)
	if response.CheckError(err, "Update link error") {
		return response.ResponseOperateFailed()
	}

	return response.ResponseOperateSuccess()
}

func (l *LinksController) AddLink(ctx *gin.Context) *response.Response {
	var link model.Link
	err := ctx.ShouldBind(&link)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	err = l.linkService.AddLink(&link)
	if response.CheckError(err, "Add link error") {
		return response.ResponseOperateFailed()
	}

	return response.ResponseOperateSuccess()
}

func (l *LinksController) Categories(ctx *gin.Context) *response.Response {
	categories, err := l.linkService.GetAllCategory()
	if response.CheckError(err, "Get categories error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(categories)
}

func (l *LinksController) DeleteCategory(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	err := l.linkService.DeleteCategory(id)
	if response.CheckError(err, "Delete category error") {
		return response.ResponseDeleteFailed()
	}

	return response.ResponseDeleteSuccess()
}

func (l *LinksController) UpdateCategory(ctx *gin.Context) *response.Response {
	var category model.LinkCategory
	err := ctx.ShouldBind(&category)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	err = l.linkService.UpdateCategory(&category)
	if response.CheckError(err, "Update category error") {
		return response.ResponseOperateFailed()
	}

	return response.ResponseOperateSuccess()
}

func (l *LinksController) AddCategory(ctx *gin.Context) *response.Response {
	var category model.LinkCategory
	err := ctx.ShouldBind(&category)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	err = l.linkService.AddCategory(&category)
	if response.CheckError(err, "Add category error") {
		return response.ResponseOperateFailed()
	}

	return response.ResponseOperateSuccess()
}