package dao

import "blog_web/model"

type LeaveMessageDao struct {
	sql []string
}

func NewLeaveMessageDao() *LeaveMessageDao {
	return &LeaveMessageDao{
		sql: []string{
			`INSERT INTO t_leave_message(nickname, email, url, content, avatar, create_time, parent_id, top_parent_id, ip)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`,

			`SELECT id, nickname, content, avatar, create_time, parent_id, top_parent_id 
			FROM t_leave_message WHERE top_parent_id = 0 AND status = 1 LIMIT ?, ?;`,

			`SELECT id, nickname, content, avatar, create_time, parent_id, top_parent_id 
			FROM t_leave_message WHERE top_parent_id = ? AND status = 1;`,

			`SELECT COUNT(*) FROM t_leave_message WHERE top_parent_id = 0 AND status = 1;`,

			`SELECT COUNT(*) FROM t_leave_message WHERE status = 1;`,

			`SELECT id, nickname, email, url, content, create_time, parent_id, top_parent_id, status, ip 
			FROM t_leave_message LIMIT ?, ?;`,

			`SELECT COUNT(*) FROM t_leave_message;`,

			`UPDATE t_leave_message SET status = ? WHERE id = ?;`,
		},
	}
}

func (l *LeaveMessageDao) InsertOne(m *model.Message) error {
	_, err := sqldb.Exec(l.sql[0], m.Nickname, m.Email, m.Url, m.Content, m.Avatar, m.CreateTime, m.ParentId, m.TopParentId, m.Ip)
	return err
}

// 获取顶级留言
func (l *LeaveMessageDao) FindTopMessage(pageStart, PageSize int) (msg []model.Message, err error) {
	err = sqldb.Select(&msg, l.sql[1], pageStart, PageSize)
	return
}

// 获取子留言
func (l *LeaveMessageDao) FindChildMessage(parentId int) (msg []model.Message, err error) {
	err = sqldb.Select(&msg, l.sql[2], parentId)
	return
}

// 查询顶级评论的数目
func (l *LeaveMessageDao) FindTopMessageCount() (count int, err error) {
	err = sqldb.Get(&count, l.sql[3])
	return
}

// 查询所有评论的数目
func (l *LeaveMessageDao) FindMessageCount() (count int, err error) {
	err = sqldb.Get(&count, l.sql[4])
	return
}

// 查询评论
func (l *LeaveMessageDao) FindLimited(pageStart, pageSize int) (messages []model.Message, err error) {
	err = sqldb.Select(&messages, l.sql[5], pageStart, pageSize)
	return
}

func (l *LeaveMessageDao) FindTotalCount() (count int, err error) {
	err = sqldb.Get(&count, l.sql[6])
	return
}

func (l *LeaveMessageDao) UpdateStatus(message *model.Message) error {
	_, err := sqldb.Exec(l.sql[7], message.Status, message.Id)
	return err
}