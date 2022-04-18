package controller

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
)

/*
* @Author: mgh
* @Date: 2022/3/1 11:07
* @Desc:
 */

type TypeListController struct {
	blogService *service.BlogService
	typeService *service.TypeService
}

func NewTypeListRouter() *TypeListController {
	return &TypeListController{
		blogService: service.NewBlogService(),
		typeService: service.NewTypeService(),
	}
}

// 博客类型页面获取所有博客类型
func (t *TypeListController) GetTypeList(ctx *gin.Context) *response.Response {
	allTypes, err := t.typeService.FindAll()
	if response.CheckError(err, "Get all types error") {
		return response.ResponseQueryFailed()
	}
	typeIds, err := t.blogService.GetAllTypes()
	if response.CheckError(err, "Get all types error") {
		return response.ResponseQueryFailed()
	}
	m := make(map[int]*model.TheType)
	for i := 0; i < len(allTypes); i++ {
		m[allTypes[i].Id] = &allTypes[i]
	}

	for _, v := range typeIds {
		m[v].Count++
	}

	return response.ResponseQuerySuccess(allTypes)
}

// 博客类型页面根据博客类型ID获取博客
func (t *TypeListController) GetBlogListByTypeid(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "8")
	typeId := utils.QueryInt(ctx, "typeId")

	blogs, count, err := t.blogService.GetBlogListByTypeId(typeId, pageNum, pageSize)
	if response.CheckError(err, "Get blogs error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blogs, count)
}
