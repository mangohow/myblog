package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
* @Author: mgh
* @Date: 2022/3/2 11:13
* @Desc:
 */

type TagController struct {
	tagService *service.TagService
}

func NewTagRouter() *TagController {
	return &TagController{
		tagService: service.NewTagService(),
	}
}

func (t *TagController) GetAllTags(ctx *gin.Context) {
	tags := t.tagService.GetAllTags()
	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, tags))
}

func (t *TagController) GetOnePageTags(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "5")
	tags, count := t.tagService.GetOnePageTags(pageNum, pageSize)
	result := utils.ResponseResult(utils.QUERY_SUCCESS, tags)
	result["total"] = count
	ctx.JSON(http.StatusOK, result)
}

func (t *TagController) CheckTagExist(ctx *gin.Context) {
	name := ctx.Query("name")
	exist := t.tagService.CheckTagExist(name)
	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, exist))
}

func (t *TagController) DeleteTag(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	ok := t.tagService.DeleteTagById(id)
	if !ok {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_SUCCESS))
}

func (t *TagController) UpdateTag(ctx *gin.Context) {
	var tag model.Tag
	if err := ctx.ShouldBind(&tag); err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ok := t.tagService.UpdateTagName(tag.Id, tag.Name)
	if !ok {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}

func (t *TagController) AddTag(ctx *gin.Context) {
	var tag model.Tag
	if err := ctx.ShouldBind(&tag); err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ok := t.tagService.AddTag(tag.Name)
	if !ok {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}
