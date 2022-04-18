package controller

import (
	"blog_web/db/service"
	"blog_web/response"
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
func (t *TimeLineController) GetTimeLinedBlogs(ctx *gin.Context) *response.Response {
	blogs, err := t.blogService.GetTimeLineBlogs()
	if response.CheckError(err, "Get timeline blogs error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blogs)
}

func (t *TimeLineController) GetGroupedBlogs(ctx *gin.Context) *response.Response {
	types, err := t.blogService.GetTypeAndBlogCount()
	if response.CheckError(err, "Get grouped blogs error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(types)
}