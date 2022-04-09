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

type BlogRouter struct {
	blogService *service.BlogService
	mottoService *service.MottoService
}

func NewBlogRouter() *BlogRouter {
	return &BlogRouter{
		blogService: service.NewBlogService(),
		mottoService: service.NewMottoService(),
	}
}

func (b *BlogRouter) SearchBlogs(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "5")
	blogTitle := ctx.Query("blogTitle")
	typeId := utils.QueryInt(ctx, "typeId")
	recommended := ctx.Query("recommended")

	blogs, count := b.blogService.GetBlogsByTitleOrTypeOrRecommend(pageNum, pageSize, blogTitle, typeId, recommended)
	result := utils.ResponseResult(utils.QUERY_SUCCESS, blogs)
	result["count"] = count
	ctx.JSON(http.StatusOK, result)
}

func (b *BlogRouter) DeleteBlog(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	ok := b.blogService.DeleteBlog(id)
	if !ok {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_SUCCESS))
}

func (b *BlogRouter) GetFullBlog(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")

	blogs, tags := b.blogService.GetFullBlog(id)
	result := utils.ResponseResult(utils.QUERY_SUCCESS, blogs)
	result["tags"] = tags
	ctx.JSON(http.StatusOK, result)
}

func (b *BlogRouter) AddBlog(ctx *gin.Context) {
	var blog model.FullBlog
	err := ctx.ShouldBind(&blog)
	if err != nil {
		utils.Logger().Warning("%v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	tm := time.Now()
	if blog.Id == 0 {
		blog.Views = 0
		blog.CreateTime = tm
	}
	blog.UpdateTime = tm

	var ok bool
	if blog.Id == 0 { // 新增博客
		ok = b.blogService.AddBlog(&blog)
	} else { // 更新博客
		ok = b.blogService.UpdateBlog(&blog)
	}
	if !ok {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}

func (b *BlogRouter) MottoList(ctx *gin.Context) {
	mottos := b.mottoService.GetAllMottoWithCreateTime()
	if mottos == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, mottos))
}

func (b *BlogRouter) AddMotto(ctx *gin.Context) {
	var motto model.Motto
	err := ctx.ShouldBind(&motto)
	if err != nil {
		utils.Logger().Warning("bind param error:%v", err)
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	motto.CreateTime = time.Now()
	err = b.mottoService.AddOne(&motto)
	if err != nil {
		utils.Logger().Warning("add motto error:%v", err)
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}

func (b *BlogRouter) DeleteMotto(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	if id == 0 {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_FAILED))
		return
	}

	err := b.mottoService.DeleteOne(uint32(id))
	if err != nil {
		utils.Logger().Warning("delete motto error:%v", err)
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_SUCCESS))
}

func (b *BlogRouter) UpdateMotto(ctx *gin.Context) {
	var motto model.Motto
	err := ctx.ShouldBind(&motto)
	if err != nil {
		utils.Logger().Warning("bind param error:%v", err)
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	err = b.mottoService.UpdateOne(&motto)
	if err != nil {
		utils.Logger().Warning("update motto error:%v", err)
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}