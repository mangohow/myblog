package controller

import (
	"blog_web/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkError(err error, msg string, args ...interface{}) bool {
	if err != nil {
		if len(args) > 0 {
			utils.Logger().Warning("%s:%v", fmt.Sprintf(msg, args...), err)
		} else {
			utils.Logger().Warning("%s:%v", msg, err)
		}
		return true
	}

	return false
}

func querySuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, utils.ResponseResult(utils.QUERY_SUCCESS, data))
}

func queryFailed(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
}

func operateFailed(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
}

func operateSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}

