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
func (h *HomeController) HomeListBlogs(ctx *gin.Context) {
	pageNum := utils.QueryInt(ctx, "pageNum")
	pageSize := utils.QueryInt(ctx, "pageSize")
	utils.Logger().Debug("pageNum:%v, pageSize:%v", pageNum, pageSize)

	if pageNum <= 0 || pageSize <= 0 {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}
	blogs, err := h.blogService.GetHomePageBlogs(pageNum, pageSize)
	if checkError(err, "Get Blogs error") {
		queryFailed(ctx)
	}
	count, err := h.blogService.GetBolgCount()
	if checkError(err, "Get Blogs error") {
		queryFailed(ctx)
	}

	result := utils.ResponseResult(utils.QUERY_SUCCESS, blogs)
	result["count"] = count
	ctx.JSON(http.StatusOK, result)
}

func (h *HomeController) GetHomePageUInfo(ctx *gin.Context) {
	user, tagn, err := h.userService.GetInfo()
	if checkError(err, "Get Userinfo error") {
		queryFailed(ctx)
		return
	}

	res := utils.ResponseResult(utils.QUERY_SUCCESS, user)
	res["tagCount"] = tagn
	ctx.JSON(http.StatusOK, res)
}

// 浏览博客详情
func (h *HomeController) GetDetailedBlog(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	blog, tags, err := h.blogService.GetDetailedBlog(id)
	if checkError(err, "Get Detailed blog error") {
		queryFailed(ctx)
		return
	}

	result := utils.ResponseResult(utils.QUERY_SUCCESS, blog)
	result["tags"] = tags
	ctx.JSON(http.StatusOK, result)
}

// 发布一条评论
func (h *HomeController) PublishComment(ctx *gin.Context) {
	var comment model.Comment
	err := ctx.ShouldBind(&comment)
	if checkError(err, "Bind param error") {
		operateFailed(ctx)
		return
	}

	err = h.commentService.Publish(&comment)
	if checkError(err, "Publish comment error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}

// 获取评论列表
func (h *HomeController) GetCommentList(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	comments, err := h.commentService.GetCommentList(id)
	if checkError(err, "Get Comment List error") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, comments)
}

// 搜索博客
func (h *HomeController) SearchBlog(ctx *gin.Context) {
	keyWord := ctx.Query("keyWord")
	blogs, err := h.blogService.GetBlogsByKeyWord(keyWord)
	if checkError(err, "Search Blogs error") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, blogs)
}

func (h *HomeController) GetBgImages(ctx *gin.Context) {
	urls, err := h.bgImageService.GetAllUrl()
	if checkError(err, "Get Background Images error") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, urls)
}

func (h *HomeController) GetMottos(ctx *gin.Context) {
	mottos, err := h.mottoService.GetAllMotto()
	if checkError(err, "Get Mottos error") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, mottos)
}

//获取最新推荐博客
func (h *HomeController) GetNewBlogs(ctx *gin.Context) {
	limit := utils.DefaultQueryInt(ctx, "countLimit", "10")
	blogs, err := h.blogService.GetNewBlogs(limit)
	if checkError(err, "Get New Blogs error") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, blogs)
}

// 获取热门博客
func (h *HomeController) GetHotBlogs(ctx *gin.Context) {
	limit := utils.DefaultQueryInt(ctx, "countLimit", "10")
	blogs, err := h.blogService.GetHotBlogs(limit)
	if checkError(err, "Get Hot Blogs error") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, blogs)
}