package router

import (
	"blog_web/db/service"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResourceLibRouter struct {
	linkService *service.LinkService
}

func NewResourceLibRouter() *ResourceLibRouter {
	return &ResourceLibRouter{
		linkService: service.NewLinkService(),
	}
}

func (r *ResourceLibRouter) LinkList(ctx *gin.Context) {
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
