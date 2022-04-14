package controller

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
func (t *TypeListController) GetTypeList(ctx *gin.Context) {
	allTypes := t.typeService.FindAll()
	typeIds := t.blogService.GetAllTypes()
	m := make(map[int]*model.TheType)
	for i := 0; i < len(allTypes); i++ {
		m[allTypes[i].Id] = &allTypes[i]
	}

	for _, v := range typeIds {
		m[v].Count++
	}

	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, allTypes))
}

// 博客类型页面根据博客类型ID获取博客
func (t *TypeListController) GetBlogListByTypeid(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "8")
	typeId := utils.QueryInt(ctx, "typeId")

	blogs, count := t.blogService.GetBlogListByTypeId(typeId, pageNum, pageSize)
	result := utils.ResponseResult(utils.QUERY_SUCCESS, blogs)
	result["count"] = count
	ctx.JSON(http.StatusOK, result)
}
