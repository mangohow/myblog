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
	tags, err := t.tagService.GetAllTags()
	if checkError(err, "Get tags error") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, tags)
}

func (t *TagController) GetOnePageTags(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "5")
	tags, count, err := t.tagService.GetOnePageTags(pageNum, pageSize)
	if checkError(err, "Get one page tag error") {
		queryFailed(ctx)
		return
	}

	result := utils.ResponseResult(utils.QUERY_SUCCESS, tags)
	result["total"] = count
	ctx.JSON(http.StatusOK, result)
}

func (t *TagController) CheckTagExist(ctx *gin.Context) {
	name := ctx.Query("name")
	exist := t.tagService.CheckTagExist(name)
	querySuccess(ctx, exist)
}

func (t *TagController) DeleteTag(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	err := t.tagService.DeleteTagById(id)
	if checkError(err, "Delete tag error") {
		deleteFailed(ctx)
		return
	}

	deleteSuccess(ctx)
}

func (t *TagController) UpdateTag(ctx *gin.Context) {
	var tag model.Tag
	if err := ctx.ShouldBind(&tag); err != nil {
		operateFailed(ctx)
		return
	}

	err := t.tagService.UpdateTagName(tag.Id, tag.Name)
	if checkError(err, "Update tag error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}

func (t *TagController) AddTag(ctx *gin.Context) {
	var tag model.Tag
	if err := ctx.ShouldBind(&tag); err != nil {
		operateFailed(ctx)
		return
	}

	err := t.tagService.AddTag(tag.Name)
	if checkError(err, "Add tag error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}
