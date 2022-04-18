package controller

import (
	"blog_web/db/service"
	"blog_web/response"
	"github.com/gin-gonic/gin"
)

type ResourceLibController struct {
	linkService *service.LinkService
}

func NewResourceLibRouter() *ResourceLibController {
	return &ResourceLibController{
		linkService: service.NewLinkService(),
	}
}

func (r *ResourceLibController) LinkList(ctx *gin.Context) *response.Response {
	links, err := r.linkService.GetAllLinks()
	if response.CheckError(err, "Get links error") {
		return response.ResponseQueryFailed()
	}
	categories, err := r.linkService.GetAllCategory()
	if response.CheckError(err, "Get categories error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(links, categories)
}
