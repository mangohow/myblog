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
	if checkError(err, "Bind param error") {
		operateFailed(ctx)
		return
	}

	// 在使用nginx作为反向代理时，需要在nginx中进行配置，添加客户端的真实IP到头部，以便web服务器获取
	msg.Ip = ctx.GetHeader("X-Forwarded-For")

	msg.Avatar = utils.RandomInt(28)
	msg.CreateTime = time.Now()
	err = l.messageService.AddMessage(&msg)
	if checkError(err, "Add message error, msg:%s", msg.Ip) {
		operateFailed(ctx)
	} else {
		operateSuccess(ctx)
	}
}

func (l *LeaveMessageController) DisplayMessage(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")

	messages, count, totalCount, err := l.messageService.GetDisplayMessage(pageNum, pageSize)
	if checkError(err, "Get messages error") {
		queryFailed(ctx)
		return
	}

	result := utils.ResponseResult(utils.QUERY_SUCCESS, messages)
	result["count"] = count
	result["totalCount"] = totalCount
	ctx.JSON(http.StatusOK, result)
}