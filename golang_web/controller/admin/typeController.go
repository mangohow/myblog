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
	types := t.typeService.FindAll()
	if types == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}
	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, types))
}

func (t *TypeController) GetOnePageTypes(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "5")
	types, count := t.typeService.GetOnePage(pageNum, pageSize)
	result := utils.ResponseResult(utils.QUERY_SUCCESS, types)
	result["total"] = count
	ctx.JSON(http.StatusOK, result)
}

func (t *TypeController) CheckTypeExist(ctx *gin.Context) {
	typename := ctx.Query("typename")
	exist := t.typeService.CheckTypeExist(typename)
	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, exist))
}

func (t *TypeController) DeleteType(ctx *gin.Context) {
	id := utils.QueryInt(ctx, "id")
	ok := t.typeService.DeleteById(id)
	if !ok {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_FAILED))
		return
	}
	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.DELETE_SUCCESS))
}

func (t *TypeController) UpdateType(ctx *gin.Context) {
	var tp model.TheType
	err := ctx.ShouldBind(&tp)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ok := t.typeService.UpdateName(tp.Id, tp.Name)
	if !ok {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}
	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}

func (t *TypeController) AddType(ctx *gin.Context) {
	var tp model.TheType
	ctx.ShouldBind(&tp)
	ok := t.typeService.AddType(tp.Name)
	if !ok {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}
	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}
