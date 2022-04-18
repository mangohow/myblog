package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
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

func (t *TagController) GetAllTags(ctx *gin.Context) *response.Response {
	tags, err := t.tagService.GetAllTags()
	if response.CheckError(err, "Get tags error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(tags)
}

func (t *TagController) GetOnePageTags(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "5")
	tags, count, err := t.tagService.GetOnePageTags(pageNum, pageSize)
	if response.CheckError(err, "Get one page tag error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(tags, count)
}

func (t *TagController) CheckTagExist(ctx *gin.Context) *response.Response {
	name := ctx.Query("name")
	exist := t.tagService.CheckTagExist(name)
	return response.ResponseQuerySuccess(exist)
}

func (t *TagController) DeleteTag(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	err := t.tagService.DeleteTagById(id)
	if response.CheckError(err, "Delete tag error") {
		return response.ResponseDeleteFailed()
	}

	return response.ResponseDeleteSuccess()
}

func (t *TagController) UpdateTag(ctx *gin.Context) *response.Response {
	var tag model.Tag
	if err := ctx.ShouldBind(&tag); err != nil {
		return response.ResponseOperateFailed()
	}

	err := t.tagService.UpdateTagName(tag.Id, tag.Name)
	if response.CheckError(err, "Update tag error") {
		return response.ResponseOperateFailed()
	}

	return response.ResponseOperateSuccess()
}

func (t *TagController) AddTag(ctx *gin.Context) *response.Response {
	var tag model.Tag
	if err := ctx.ShouldBind(&tag); err != nil {
		return response.ResponseOperateFailed()
	}

	err := t.tagService.AddTag(tag.Name)
	if response.CheckError(err, "Add tag error") {
		return response.ResponseOperateFailed()
	}

	return response.ResponseOperateSuccess()
}
