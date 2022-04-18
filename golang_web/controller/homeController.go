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
* @Date: 2022/2/28 19:04
* @Desc: 博客主页对应的Routers
 */

type HomeController struct {
	blogService    *service.BlogService
	commentService *service.CommentService
	bgImageService *service.BGImageService
	mottoService *service.MottoService
	userService *service.UserService
}

func NewHomeRouter() *HomeController {
	return &HomeController{
		blogService:    service.NewBlogService(),
		commentService: service.NewCommentService(),
		bgImageService: service.NewBGImageService(),
		mottoService: service.NewMottoService(),
		userService: service.NewUserService(),
	}
}

// 主页博客展示
func (h *HomeController) HomeListBlogs(ctx *gin.Context) *response.Response {
	pageNum := utils.QueryInt(ctx, "pageNum")
	pageSize := utils.QueryInt(ctx, "pageSize")
	utils.Logger().Debug("pageNum:%v, pageSize:%v", pageNum, pageSize)

	if pageNum <= 0 || pageSize <= 0 {
		return response.ResponseQueryFailed()
	}
	blogs, err := h.blogService.GetHomePageBlogs(pageNum, pageSize)
	if response.CheckError(err, "Get Blogs error") {
		return response.ResponseQueryFailed()
	}
	count, err := h.blogService.GetBolgCount()
	if response.CheckError(err, "Get Blogs error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blogs, count)
}

func (h *HomeController) GetHomePageUInfo(ctx *gin.Context) *response.Response {
	user, tagn, err := h.userService.GetInfo()
	if response.CheckError(err, "Get Userinfo error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(user, tagn)
}

// 浏览博客详情
func (h *HomeController) GetDetailedBlog(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	blog, tags, err := h.blogService.GetDetailedBlog(id)
	if response.CheckError(err, "Get Detailed blog error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blog, tags)
}

// 发布一条评论
func (h *HomeController) PublishComment(ctx *gin.Context) *response.Response {
	var comment model.Comment
	err := ctx.ShouldBind(&comment)
	if response.CheckError(err, "Bind param error") {
		return response.ResponseOperateFailed()
	}

	err = h.commentService.Publish(&comment)
	if response.CheckError(err, "Publish comment error") {
		return response.ResponseOperateFailed()
	}

	return response.ResponseOperateSuccess()
}

// 获取评论列表
func (h *HomeController) GetCommentList(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	comments, err := h.commentService.GetCommentList(id)
	if response.CheckError(err, "Get Comment List error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(comments)
}

// 搜索博客
func (h *HomeController) SearchBlog(ctx *gin.Context) *response.Response {
	keyWord := ctx.Query("keyWord")
	blogs, err := h.blogService.GetBlogsByKeyWord(keyWord)
	if response.CheckError(err, "Search Blogs error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blogs)
}

func (h *HomeController) GetBgImages(ctx *gin.Context) *response.Response {
	urls, err := h.bgImageService.GetAllUrl()
	if response.CheckError(err, "Get Background Images error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(urls)
}

func (h *HomeController) GetMottos(ctx *gin.Context) *response.Response {
	mottos, err := h.mottoService.GetAllMotto()
	if response.CheckError(err, "Get Mottos error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(mottos)
}

//获取最新推荐博客
func (h *HomeController) GetNewBlogs(ctx *gin.Context) *response.Response {
	limit := utils.DefaultQueryInt(ctx, "countLimit", "10")
	blogs, err := h.blogService.GetNewBlogs(limit)
	if response.CheckError(err, "Get New Blogs error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blogs)
}

// 获取热门博客
func (h *HomeController) GetHotBlogs(ctx *gin.Context) *response.Response {
	limit := utils.DefaultQueryInt(ctx, "countLimit", "10")
	blogs, err := h.blogService.GetHotBlogs(limit)
	if response.CheckError(err, "Get Hot Blogs error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blogs)
}