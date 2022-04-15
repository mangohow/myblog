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

	links, err := l.linkService.GetLimitedLinks(pageNum, pageSize)
	if checkError(err, "Get links error") {
		queryFailed(ctx)
		return
	}
	categories, err := l.linkService.GetAllCategory()
	if checkError(err, "Get categories error") {
		queryFailed(ctx)
		return
	}

	count, _ := l.linkService.GetLinkCount()
	result := utils.ResponseResult(utils.QUERY_SUCCESS, links)
	result["categories"] = categories
	result["count"] = count
	ctx.JSON(http.StatusOK, result)
}

func (l *LinksController) DeleteLink(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	err := l.linkService.DeleteLink(id)
	if checkError(err, "Delete link error") {
		deleteFailed(ctx)
		return
	}

	deleteSuccess(ctx)
}

func (l *LinksController) UpdateLink(ctx *gin.Context) {
	var link model.Link
	err := ctx.ShouldBind(&link)
	if checkError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = l.linkService.UpdateLink(&link)
	if checkError(err, "Update link error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}

func (l *LinksController) AddLink(ctx *gin.Context) {
	var link model.Link
	err := ctx.ShouldBind(&link)
	if checkError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = l.linkService.AddLink(&link)
	if checkError(err, "Add link error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}

func (l *LinksController) Categories(ctx *gin.Context) {
	categories, err := l.linkService.GetAllCategory()
	if checkError(err, "Get categories error") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, categories)
}

func (l *LinksController) DeleteCategory(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	err := l.linkService.DeleteCategory(id)
	if checkError(err, "Delete category error") {
		deleteFailed(ctx)
		return
	}

	deleteSuccess(ctx)
}

func (l *LinksController) UpdateCategory(ctx *gin.Context) {
	var category model.LinkCategory
	err := ctx.ShouldBind(&category)
	if checkError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = l.linkService.UpdateCategory(&category)
	if checkError(err, "Update category error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}

func (l *LinksController) AddCategory(ctx *gin.Context) {
	var category model.LinkCategory
	err := ctx.ShouldBind(&category)
	if checkError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = l.linkService.AddCategory(&category)
	if checkError(err, "Add category error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}