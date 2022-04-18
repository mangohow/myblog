package controller

import (
	"blog_web/db/service"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
)

/*
* @Author: mgh
* @Date: 2022/3/1 11:51
* @Desc:
 */

type TagListController struct {
	tagService  *service.TagService
	blogService *service.BlogService
}

func NewTagListRouter() *TagListController {
	return &TagListController{
		tagService:  service.NewTagService(),
		blogService: service.NewBlogService(),
	}
}

// 标签页获取所有的标签
func (t *TagListController) GetTagList(ctx *gin.Context) *response.Response {
	tags, err := t.tagService.GetAllTags()
	if response.CheckError(err, "Get tags error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(tags)
}

// 标签页根据标签ID获取博客
func (t *TagListController) GetBlogListByTagId(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "8")
	id := utils.QueryInt(ctx, "tagId")

	blogs, i, err := t.blogService.GetBlogsByTagId(id, pageNum, pageSize)
	if response.CheckError(err, "get blogs error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blogs, i)
}
