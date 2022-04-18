package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageController struct {
	msgService *service.LeaveMessageService
}

func NewMessageRouter() *MessageController {
	return &MessageController{
		msgService: service.NewLeaveMessageService(),
	}
}

func (m *MessageController) MessageList(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")
	messages, count, err := m.msgService.GetPageMessage(pageNum, pageSize)
	if response.CheckError(err, "Get messages error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(messages, count)
}

func (m *MessageController) UpdateStatus(ctx *gin.Context) *response.Response {
	var msg model.Message
	err := ctx.ShouldBind(&msg)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	err = m.msgService.UpdateMsgStatus(&msg)
	if response.CheckError(err, "Update message status error") {
		return response.ResponseOperateFailed()
	}

	return response.ResponseOperateSuccess()
}

