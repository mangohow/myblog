package controller

import (
	"blog_web/db/service"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
	blogs := t.blogService.GetTimeLineBlogs()
	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, blogs))
}

func (t *TimeLineController) GetGroupedBlogs(ctx *gin.Context) {
	types := t.blogService.GetTypeAndBlogCount()
	if types == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, types))
}