package service

import (
	"blog_web/db/dao"
	"blog_web/model"
	"blog_web/utils"
)

type LeaveMessageService struct {
	messageDao *dao.LeaveMessageDao
}

func NewLeaveMessageService() *LeaveMessageService {
	return &LeaveMessageService{
		messageDao: dao.NewLeaveMessageDao(),
	}
}

func (l *LeaveMessageService) AddMessage(m *model.Message) error {
	return l.messageDao.InsertOne(m)
}

// 获取留言页展示的留言
func (l *LeaveMessageService) GetDisplayMessage(pageNum, pageSize int) ([]model.Message, int, int, error) {
	pageStart := (pageNum - 1) * pageSize
	messages, err := l.messageDao.FindTopMessage(pageStart, pageSize)
	if err != nil {
		utils.Logger().Warning("%v", err)
		return nil, 0, 0, err
	}
	totalCount, err := l.messageDao.FindMessageCount()
	count, _ := l.messageDao.FindTopMessageCount()
	for _, v := range messages {
		msg, err := l.messageDao.FindChildMessage(v.Id)
		if err == nil {
			messages = append(messages, msg...)
		}
	}

	return messages, count, totalCount, nil
}

func (l *LeaveMessageService) GetPageMessage(pageNum, pageSize int) ([]model.Message, int, error) {
	pageStart := (pageNum - 1) * pageSize
	messages, err := l.messageDao.FindLimited(pageStart, pageSize)
	if err != nil {
		utils.Logger().Warning("get limited message error:%v", err)
		return nil, 0, err
	}
	count, _ := l.messageDao.FindTotalCount()

	return messages, count, nil
}

func (l *LeaveMessageService) UpdateMsgStatus(message *model.Message) error {
	return l.messageDao.UpdateStatus(message)
}