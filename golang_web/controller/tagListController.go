package controller

import (
	"blog_web/db/service"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
func (t *TagListController) GetTagList(ctx *gin.Context) {
	tags, err := t.tagService.GetAllTags()
	if checkError(err, "Get tags error") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, tags)
}

// 标签页根据标签ID获取博客
func (t *TagListController) GetBlogListByTagId(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "8")
	id := utils.QueryInt(ctx, "tagId")

	blogs, i, err := t.blogService.GetBlogsByTagId(id, pageNum, pageSize)
	if checkError(err, "get blogs error") {
		queryFailed(ctx)
		return
	}

	result := utils.ResponseResult(utils.QUERY_SUCCESS, blogs)
	result["count"] = i
	ctx.JSON(http.StatusOK, result)
}
