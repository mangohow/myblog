package utils

import "github.com/gin-gonic/gin"

/*
* @Author: mgh
* @Date: 2022/2/24 18:45
* @Desc:
 */

const (
	QUERY_FAILED uint32 = iota
	QUERY_SUCCESS

	OPERATE_FAILED = iota + 98
	OPERATE_SUCCESS

	LOGIN_FAILED = iota + 196
	LOGIN_SUCCESS

	UNAUTHORIZED = iota + 294

	DELETE_FAILED = iota + 393
	DELETE_SUCCESS
)

var Descs = map[uint32]string{
	QUERY_SUCCESS:   "查询成功",
	QUERY_FAILED:    "查询失败",
	OPERATE_FAILED:  "操作失败",
	OPERATE_SUCCESS: "操作成功",
	LOGIN_FAILED:    "登录失败",
	LOGIN_SUCCESS:   "登录成功",
	UNAUTHORIZED:    "未获得授权",
	DELETE_FAILED:   "删除失败",
	DELETE_SUCCESS:  "删除成功",
}

func ResponseResult(code uint32, data interface{}) gin.H {
	return gin.H{
		"status":  code,
		"message": Descs[code],
		"data":    data,
	}
}

func ResponseWithoutData(code uint32) gin.H {
	return gin.H{
		"status":  code,
		"message": Descs[code],
	}
}
