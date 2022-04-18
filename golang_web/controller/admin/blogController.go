package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
* @Author: mgh
* @Date: 2022/3/2 10:53
* @Desc:
 */

type BlogController struct {
	blogService *service.BlogService
	mottoService *service.MottoService
}

func NewBlogRouter() *BlogController {
	return &BlogController{
		blogService: service.NewBlogService(),
		mottoService: service.NewMottoService(),
	}
}

func (b *BlogController) SearchBlogs(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "5")
	blogTitle := ctx.Query("blogTitle")
	typeId := utils.QueryInt(ctx, "typeId")
	recommended := ctx.Query("recommended")

	blogs, count, err := b.blogService.GetBlogsByTitleOrTypeOrRecommend(pageNum, pageSize, blogTitle, typeId, recommended)
	if response.CheckError(err, "Search blogs error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blogs, count)
}

func (b *BlogController) DeleteBlog(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	err := b.blogService.DeleteBlog(id)
	if response.CheckError(err, "Delete blog error") {
		return response.ResponseDeleteFailed()
	}

	return response.ResponseDeleteSuccess()
}

func (b *BlogController) GetFullBlog(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	blogs, tags, err := b.blogService.GetFullBlog(id)
	if response.CheckError(err, "Get Full blog error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blogs, tags)
}

func (b *BlogController) AddBlog(ctx *gin.Context) *response.Response {
	var blog model.FullBlog
	err := ctx.ShouldBind(&blog)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	tm := time.Now()
	if blog.Id == 0 {
		blog.Views = 0
		blog.CreateTime = tm
	}
	blog.UpdateTime = tm

	if blog.Id == 0 { // 新增博客
		err = b.blogService.AddBlog(&blog)
		if response.CheckError(err, "Add blog error") {
			return response.ResponseOperateFailed()
		}
	} else { // 更新博客
		err = b.blogService.UpdateBlog(&blog)
		if response.CheckError(err, "Update blog error") {
			return response.ResponseOperateFailed()
		}
	}

	return response.ResponseOperateSuccess()
}

func (b *BlogController) MottoList(ctx *gin.Context) *response.Response {
	mottos, err := b.mottoService.GetAllMottoWithCreateTime()
	if response.CheckError(err, "Get mottos error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(mottos)
}

func (b *BlogController) AddMotto(ctx *gin.Context) *response.Response {
	var motto model.Motto
	err := ctx.ShouldBind(&motto)
	if response.CheckError(err, "Bind param error") {
		return response.ResponseOperateFailed()
	}

	motto.CreateTime = time.Now()
	err = b.mottoService.AddOne(&motto)
	if response.CheckError(err, "Add motto error") {
		return response.ResponseOperateFailed()
	}

	return response.ResponseOperateSuccess()
}

func (b *BlogController) DeleteMotto(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	if id == 0 {
		return response.ResponseDeleteFailed()
	}

	err := b.mottoService.DeleteOne(uint32(id))
	if response.CheckError(err, "Delete motto error") {
		return response.ResponseDeleteFailed()
	}

	return response.ResponseDeleteSuccess()
}

func (b *BlogController) UpdateMotto(ctx *gin.Context) *response.Response {
	var motto model.Motto
	err := ctx.ShouldBind(&motto)
	if response.CheckError(err, "Bind param error") {
		return response.ResponseOperateFailed()
	}

	err = b.mottoService.UpdateOne(&motto)
	if response.CheckError(err, "Update motto error") {
		return response.ResponseOperateFailed()
	}

	return response.ResponseOperateSuccess()
}