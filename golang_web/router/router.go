package router

import (
	"blog_web/response"
	"github.com/gin-gonic/gin"
)

type Handler func(ctx *gin.Context) *response.Response

type Engine struct {
	*gin.Engine
}

func NewEngine(e *gin.Engine) *Engine {
	return &Engine{e}
}

type RouterGroup struct {
	*gin.RouterGroup
}

func NewRouterGroup(r *gin.RouterGroup) *RouterGroup {
	return &RouterGroup{r}
}

func (e *Engine) Get(relativePath string, handler Handler) {
	e.GET(relativePath, func(ctx *gin.Context) {
		r := handler(ctx)
		if r != nil {
			ctx.JSON(r.HttpStatus, &r.R)
		}

		response.PutResponse(r)
	})
}

func (e *Engine) Post(relativePath string, handler Handler) {
	e.POST(relativePath, func(ctx *gin.Context) {
		r := handler(ctx)
		if r != nil {
			ctx.JSON(r.HttpStatus, &r.R)
		}

		response.PutResponse(r)
	})
}

func (e *Engine) Put(relativePath string, handler Handler) {
	e.PUT(relativePath, func(ctx *gin.Context) {
		r := handler(ctx)
		if r != nil {
			ctx.JSON(r.HttpStatus, &r.R)
		}

		response.PutResponse(r)
	})
}

func (e *Engine) Delete(relativePath string, handler Handler) {
	e.DELETE(relativePath, func(ctx *gin.Context) {
		r := handler(ctx)
		if r != nil {
			ctx.JSON(r.HttpStatus, &r.R)
		}

		response.PutResponse(r)
	})
}



func (r *RouterGroup) Get(relativePath string, handler Handler) {
	r.GET(relativePath, func(ctx *gin.Context) {
		r := handler(ctx)
		if r != nil {
			ctx.JSON(r.HttpStatus, &r.R)
		}

		response.PutResponse(r)
	})
}

func (r *RouterGroup) Post(relativePath string, handler Handler) {
	r.POST(relativePath, func(ctx *gin.Context) {
		r := handler(ctx)
		if r != nil {
			ctx.JSON(r.HttpStatus, &r.R)
		}

		response.PutResponse(r)
	})
}

func (r *RouterGroup) Put(relativePath string, handler Handler) {
	r.PUT(relativePath, func(ctx *gin.Context) {
		r := handler(ctx)
		if r != nil {
			ctx.JSON(r.HttpStatus, &r.R)
		}

		response.PutResponse(r)
	})
}

func (r *RouterGroup) Delete(relativePath string, handler Handler) {
	r.DELETE(relativePath, func(ctx *gin.Context) {
		r := handler(ctx)
		if r != nil {
			ctx.JSON(r.HttpStatus, &r.R)
		}

		response.PutResponse(r)
	})
}
