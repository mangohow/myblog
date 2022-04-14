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
	blogs := h.blogService.GetHomePageBlogs(pageNum, pageSize)
	count := h.blogService.GetBolgCount()
	result := utils.ResponseResult(utils.QUERY_SUCCESS, blogs)
	result["count"] = count
	ctx.JSON(http.StatusOK, result)
}

func (h *HomeController) GetHomePageUInfo(ctx *gin.Context) {
	user, tagn := h.userService.GetInfo()
	if user == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}

	res := utils.ResponseResult(utils.QUERY_SUCCESS, user)
	res["tagCount"] = tagn
	ctx.JSON(http.StatusOK, res)
}

// 浏览博客详情
func (h *HomeController) GetDetailedBlog(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	blog, tags := h.blogService.GetDetailedBlog(id)
	if blog == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
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
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}
	err = h.commentService.Publish(&comment)
	if err != nil {
		utils.Logger().Warning("%v", err)
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}

// 获取评论列表
func (h *HomeController) GetCommentList(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	comments := h.commentService.GetCommentList(id)
	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, comments))
}

// 搜索博客
func (h *HomeController) SearchBlog(ctx *gin.Context) {
	keyWord := ctx.Query("keyWord")
	blogs := h.blogService.GetBlogsByKeyWord(keyWord)
	result := utils.ResponseResult(utils.QUERY_SUCCESS, blogs)
	ctx.JSON(http.StatusOK, result)
}

func (h *HomeController) GetBgImages(ctx *gin.Context) {
	urls := h.bgImageService.GetAllUrl()
	if urls == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}
	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, urls))
}

func (h *HomeController) GetMottos(ctx *gin.Context) {
	mottos := h.mottoService.GetAllMotto()
	if mottos == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}
	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, mottos))
}

//获取最新推荐博客
func (h *HomeController) GetNewBlogs(ctx *gin.Context) {
	limit := utils.DefaultQueryInt(ctx, "countLimit", "10")
	blogs := h.blogService.GetNewBlogs(limit)
	if blogs == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, blogs))
}

// 获取热门博客
func (h *HomeController) GetHotBlogs(ctx *gin.Context) {
	limit := utils.DefaultQueryInt(ctx, "countLimit", "10")
	blogs := h.blogService.GetHotBlogs(limit)
	if blogs == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, blogs))
}