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
* @Date: 2022/3/2 11:11
* @Desc:
 */

type TypeController struct {
	typeService *service.TypeService
}

func NewTypeRouter() *TypeController {
	return &TypeController{
		typeService: service.NewTypeService(),
	}
}

func (t *TypeController) GetAllTypes(ctx *gin.Context) {
	types, err := t.typeService.FindAll()
	if checkError(err, "Find types error") {
		queryFailed(ctx)
		return
	}

	querySuccess(ctx, types)
}

func (t *TypeController) GetOnePageTypes(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "5")
	types, count, err := t.typeService.GetOnePage(pageNum, pageSize)
	if checkError(err, "Get one page types error") {
		queryFailed(ctx)
		return
	}

	result := utils.ResponseResult(utils.QUERY_SUCCESS, types)
	result["total"] = count
	ctx.JSON(http.StatusOK, result)
}

func (t *TypeController) CheckTypeExist(ctx *gin.Context) {
	typename := ctx.Query("typename")
	exist := t.typeService.CheckTypeExist(typename)
	querySuccess(ctx, exist)
}

func (t *TypeController) DeleteType(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	err := t.typeService.DeleteById(id)
	if checkError(err, "Delete type error") {
		deleteFailed(ctx)
		return
	}

	deleteSuccess(ctx)
}

func (t *TypeController) UpdateType(ctx *gin.Context) {
	var tp model.TheType
	err := ctx.ShouldBind(&tp)
	if checkError(err, "Bind param error") {
		operateFailed(ctx)
		return
	}

	err = t.typeService.UpdateName(tp.Id, tp.Name)
	if checkError(err, "Update type error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}

func (t *TypeController) AddType(ctx *gin.Context) {
	var tp model.TheType
	ctx.ShouldBind(&tp)
	err := t.typeService.AddType(tp.Name)
	if checkError(err, "Add type error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}
