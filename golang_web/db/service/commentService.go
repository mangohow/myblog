package service

import (
	"blog_web/db/dao"
	"blog_web/model"
	"time"
)

/*
* @Author: mgh
* @Date: 2022/3/1 10:13
* @Desc:
 */

type CommentService struct {
	commentDao *dao.CommentDao
}

func NewCommentService() *CommentService {
	return &CommentService{
		commentDao: dao.NewCommentDao(),
	}
}

// 根据博客ID获取所有对应的评论
func (c *CommentService) GetCommentList(id int) ([]model.Comment, error) {
	return c.commentDao.FindByBlogId(id)
}

// 发布一条评论
func (c *CommentService) Publish(comment *model.Comment) error {
	comment.CreateTime = time.Now()
	return c.commentDao.Insert(comment)
}
