package controller

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LeaveMessageController struct {
	messageService *service.LeaveMessageService
}

func NewLeaveMessageRouter() *LeaveMessageController {
	return &LeaveMessageController{
		messageService: service.NewLeaveMessageService(),
	}
}


func (l *LeaveMessageController) LeaveMessage(ctx *gin.Context) {
	var msg model.Message
	err := ctx.ShouldBind(&msg)
	if err != nil {
		utils.Logger().Debug("%v", err)
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	// 在使用nginx作为反向代理时，需要在nginx中进行配置，添加客户端的真实IP到头部，以便web服务器获取
	msg.Ip = ctx.GetHeader("X-Forwarded-For")
	utils.Logger().Debug("%v,%v", msg.Ip, ctx.GetHeader("Host"))

	msg.Avatar = utils.RandomInt(28)
	msg.CreateTime = time.Now()
	err = l.messageService.AddMessage(&msg)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
	} else {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
	}
}

func (l *LeaveMessageController) DisplayMessage(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")

	messages, count, totalCount := l.messageService.GetDisplayMessage(pageNum, pageSize)
	if messages == nil {
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.QUERY_FAILED))
		return
	}

	result := utils.ResponseResult(utils.QUERY_SUCCESS, messages)
	result["count"] = count
	result["totalCount"] = totalCount
	ctx.JSON(http.StatusOK, result)
}