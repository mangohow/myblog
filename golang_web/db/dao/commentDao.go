package dao

import (
	"blog_web/model"
)

/*
* @Author: mgh
* @Date: 2022/3/1 10:15
* @Desc: 操作t_comment表的dao
 */

type CommentDao struct {
	sql []string
}

func NewCommentDao() *CommentDao {
	return &CommentDao{
		sql: []string{
			`SELECT * FROM t_comment WHERE blog_id = ?;`,

			`INSERT INTO t_comment (nickname, email, content, avatar, create_time, blog_id)
        	VALUES (?, ?, ?, ?, ?, ?)`,
		},
	}
}

// 根据博客Id获取对应的评论
func (c *CommentDao) FindByBlogId(id int) (comments []model.Comment, err error) {
	err = sqldb.Select(&comments, c.sql[0], id)
	return
}

// 添加一条评论
func (c *CommentDao) Insert(comment *model.Comment) error {
	_, err := sqldb.Exec(c.sql[1], comment.Nickname, comment.Email,
		comment.Content, comment.Avatar, comment.CreateTime, comment.BlogId)
	return err
}
