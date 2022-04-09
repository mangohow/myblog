package router

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

type TagListRouter struct {
	tagService  *service.TagService
	blogService *service.BlogService
}

func NewTagListRouter() *TagListRouter {
	return &TagListRouter{
		tagService:  service.NewTagService(),
		blogService: service.NewBlogService(),
	}
}

// 标签页获取所有的标签
func (t *TagListRouter) GetTagList(ctx *gin.Context) {
	tags := t.tagService.GetAllTags()
	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, tags))
}

// 标签页根据标签ID获取博客
func (t *TagListRouter) GetBlogListByTagId(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "8")
	id := utils.QueryInt(ctx, "tagId")

	blogs, i := t.blogService.GetBlogsByTagId(id, pageNum, pageSize)

	result := utils.ResponseResult(utils.QUERY_SUCCESS, blogs)
	result["count"] = i
	ctx.JSON(http.StatusOK, result)
}
