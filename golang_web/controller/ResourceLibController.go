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
	links := r.linkService.GetAllLinks()
	categories := r.linkService.GetAllCategory()
	if links == nil || categories == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}
	result := utils.ResponseResult(utils.QUERY_SUCCESS, links)
	result["categories"] = categories
	ctx.JSON(http.StatusOK, result)
}
