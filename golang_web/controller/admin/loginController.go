package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
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

type LoginController struct {
	userService *service.UserService
}

func NewLoginRouter() *LoginController {
	return &LoginController{
		userService: service.NewUserService(),
	}
}

// 博客后台登录的router
func (l *LoginController) Login(ctx *gin.Context) *response.Response {
	var u model.User
	err := ctx.ShouldBind(&u)
	if response.CheckError(err, "Bind param error") {
		return response.NewResponseOkND(response.LoginFailed)
	}

	if u.Username == "" || u.Password == "" {
		return response.NewResponseOkND(response.LoginFailed)
	}

	user, err := l.userService.CheckUser(u.Username, u.Password)
	if response.CheckError(err, "Username or Password incorrect, IP:%s", ctx.GetHeader("X-Forwarded-For")) {
		return response.NewResponseOkND(response.LoginFailed)
	}

	token, err := utils.CreateToken(uint32(user.Id), user.Username, time.Hour * 24)
	if response.CheckError(err, "Generate token error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	return response.NewResponseOk(response.LoginSuccess, token, user.Id)
}

// 用户鉴权middleware
func LoginAuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			utils.Logger().Warning("未获得授权, ip:%s", ctx.Request.RemoteAddr)
			ctx.JSON(http.StatusOK, &(response.NewResponseOkND(response.Unauthorized).R))
			ctx.Abort()
			return
		}

		if _, _, ok := utils.VerifyToken(token); !ok {
			ctx.JSON(http.StatusOK, &(response.NewResponseOkND(response.Unauthorized).R))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
