package admin

import (
	"blog_web/db/service"
	"blog_web/model"
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

func (b *BlogController) SearchBlogs(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "5")
	blogTitle := ctx.Query("blogTitle")
	typeId := utils.QueryInt(ctx, "typeId")
	recommended := ctx.Query("recommended")

	blogs, count, err := b.blogService.GetBlogsByTitleOrTypeOrRecommend(pageNum, pageSize, blogTitle, typeId, recommended)
	if checkError(err, "Search blogs error") {
		queryFailed(ctx)
		return
	}
	result := utils.ResponseResult(utils.QUERY_SUCCESS, blogs)
	result["count"] = count
	ctx.JSON(http.StatusOK, result)
}

func (b *BlogController) DeleteBlog(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	err := b.blogService.DeleteBlog(id)
	if checkError(err, "Delete blog error") {
		deleteFailed(ctx)
		return
	}

	deleteSuccess(ctx)
}

func (b *BlogController) GetFullBlog(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	blogs, tags, err := b.blogService.GetFullBlog(id)
	if checkError(err, "Get Full blog error") {
		queryFailed(ctx)
		return
	}

	result := utils.ResponseResult(utils.QUERY_SUCCESS, blogs)
	result["tags"] = tags
	ctx.JSON(http.StatusOK, result)
}

func (b *BlogController) AddBlog(ctx *gin.Context) {
	var blog model.FullBlog
	err := ctx.ShouldBind(&blog)
	if checkError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	tm := time.Now()
	if blog.Id == 0 {
		blog.Views = 0
		blog.CreateTime = tm
	}
	blog.UpdateTime = tm

	if blog.Id == 0 { // 新增博客
		err = b.blogService.AddBlog(&blog)
		if checkError(err, "Add blog error") {
			operateFailed(ctx)
			return
		}
	} else { // 更新博客
		err = b.blogService.UpdateBlog(&blog)
		if checkError(err, "Update blog error") {
			operateFailed(ctx)
			return
		}
	}

	operateSuccess(ctx)
}

func (b *BlogController) MottoList(ctx *gin.Context) {
	mottos, err := b.mottoService.GetAllMottoWithCreateTime()
	if checkError(err, "Get mottos error") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, mottos)
}

func (b *BlogController) AddMotto(ctx *gin.Context) {
	var motto model.Motto
	err := ctx.ShouldBind(&motto)
	if checkError(err, "Bind param error") {
		operateFailed(ctx)
		return
	}

	motto.CreateTime = time.Now()
	err = b.mottoService.AddOne(&motto)
	if checkError(err, "Add motto error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}

func (b *BlogController) DeleteMotto(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	if id == 0 {
		deleteFailed(ctx)
		return
	}

	err := b.mottoService.DeleteOne(uint32(id))
	if checkError(err, "Delete motto error") {
		deleteFailed(ctx)
		return
	}

	deleteSuccess(ctx)
}

func (b *BlogController) UpdateMotto(ctx *gin.Context) {
	var motto model.Motto
	err := ctx.ShouldBind(&motto)
	if checkError(err, "Bind param error") {
		operateFailed(ctx)
		return
	}

	err = b.mottoService.UpdateOne(&motto)
	if checkError(err, "Update motto error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}