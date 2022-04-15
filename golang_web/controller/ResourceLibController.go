package controller

import (
	"blog_web/db/service"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResourceLibController struct {
	linkService *service.LinkService
}

func NewResourceLibRouter() *ResourceLibController {
	return &ResourceLibController{
		linkService: service.NewLinkService(),
	}
}

func (r *ResourceLibController) LinkList(ctx *gin.Context) {
	links, err := r.linkService.GetAllLinks()
	if checkError(err, "Get links error") {
		queryFailed(ctx)
		return
	}
	categories, err := r.linkService.GetAllCategory()
	if checkError(err, "Get categories error") {
		queryFailed(ctx)
		return
	}

	result := utils.ResponseResult(utils.QUERY_SUCCESS, links)
	result["categories"] = categories
	ctx.JSON(http.StatusOK, result)
}
