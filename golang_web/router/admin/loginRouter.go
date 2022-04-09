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
* @Date: 2022/3/2 11:10
* @Desc:
 */

type LoginRouter struct {
	userService *service.UserService
}

func NewLoginRouter() *LoginRouter {
	return &LoginRouter{
		userService: service.NewUserService(),
	}
}

// 博客后台登录的router
func (l *LoginRouter) Login(ctx *gin.Context) {
	var u model.User
	err := ctx.ShouldBind(&u)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.LOGIN_FAILED))
		return
	}

	if u.Username == "" || u.Password == "" {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.LOGIN_FAILED))
		return
	}

	user := l.userService.CheckUser(u.Username, u.Password)
	if user == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.LOGIN_FAILED))
		return
	}
	token, err := utils.CreateToken(uint32(user.Id), user.Username, time.Hour * 24)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.LOGIN_FAILED))
		return
	}
	result := utils.ResponseWithoutData(utils.LOGIN_SUCCESS)
	result["token"] = token
	result["id"] = user.Id
	ctx.JSON(http.StatusOK, result)

	return
}

// 用户鉴权middleware
func LoginAuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			utils.Logger().Warning("未获得授权, ip:%s", ctx.Request.RemoteAddr)
			ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.UNAUTHORIZED))
			ctx.Abort()
			return
		}

		if _, _, ok := utils.VerifyToken(token); !ok {
			ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.UNAUTHORIZED))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
