package model

import "time"

/*
* @Author: mgh
* @Date: 2022/2/28 18:58
* @Desc:
 */

type Comment struct {
	Id         int       `db:"id" json:"id"`
	Nickname   string    `db:"nickname" json:"nickname"`
	Email      string    `db:"email" json:"email"`
	Content    string    `db:"content" json:"content"`
	Avatar     string    `db:"avatar" json:"avatar"`
	CreateTime time.Time `db:"create_time" json:"createTime"`
	BlogId     int       `db:"blog_id" json:"blogId"`
}
