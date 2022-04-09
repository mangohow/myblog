package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageRouter struct {
	msgService *service.LeaveMessageService
}

func NewMessageRouter() *MessageRouter {
	return &MessageRouter{
		msgService: service.NewLeaveMessageService(),
	}
}

func (m *MessageRouter) MessageList(ctx *gin.Context) {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")
	messages, count := m.msgService.GetPageMessage(pageNum, pageSize)
	result := utils.ResponseResult(utils.QUERY_SUCCESS, messages)
	result["count"] = count
	ctx.JSON(http.StatusOK, result)
}

func (m *MessageRouter) UpdateStatus(ctx *gin.Context) {
	var msg model.Message
	err := ctx.ShouldBind(&msg)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = m.msgService.UpdateMsgStatus(&msg)
	if err != nil {
		utils.Logger().Warning("update message status error:%v", err)
		ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_FAILED))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseWithoutData(utils.OPERATE_SUCCESS))
}

