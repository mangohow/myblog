package admin

import (
	"blog_web/db/service"
	"blog_web/model"
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

func (m *MessageController) MessageList(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")
	messages, count, err := m.msgService.GetPageMessage(pageNum, pageSize)
	if checkError(err, "Get messages error") {
		queryFailed(ctx)
		return
	}

	result := utils.ResponseResult(utils.QUERY_SUCCESS, messages)
	result["count"] = count
	ctx.JSON(http.StatusOK, result)
}

func (m *MessageController) UpdateStatus(ctx *gin.Context) {
	var msg model.Message
	err := ctx.ShouldBind(&msg)
	if checkError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = m.msgService.UpdateMsgStatus(&msg)
	if checkError(err, "Update message status error") {
		operateFailed(ctx)
		return
	}

	operateSuccess(ctx)
}

