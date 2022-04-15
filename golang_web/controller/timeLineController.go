package controller

import (
	"blog_web/db/service"
	"github.com/gin-gonic/gin"
)

/*
* @Author: mgh
* @Date: 2022/3/1 17:23
* @Desc:
 */

type TimeLineController struct {
	blogService *service.BlogService
}

func NewTimeLineRouter() *TimeLineController {
	return &TimeLineController{
		blogService: service.NewBlogService(),
	}
}

// 时间线页面获取博客
func (t *TimeLineController) GetTimeLinedBlogs(ctx *gin.Context) {
	blogs, err := t.blogService.GetTimeLineBlogs()
	if checkError(err, "Get timeline blogs error") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, blogs)
}

func (t *TimeLineController) GetGroupedBlogs(ctx *gin.Context) {
	types, err := t.blogService.GetTypeAndBlogCount()
	if checkError(err, "Get grouped blogs error") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, types)
}