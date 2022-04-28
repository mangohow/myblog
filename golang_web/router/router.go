package router

import (
	"blog_web/response"
	"github.com/gin-gonic/gin"
)

type Handler func(ctx *gin.Context) *response.Response

// 装饰器
func Decorate(h Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := h(ctx)
		if r != nil {
			ctx.JSON(r.HttpStatus, &r.R)
		}

		response.PutResponse(r)
	}
}
