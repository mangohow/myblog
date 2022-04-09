package model

import "time"

/*
* @Author: mgh
* @Date: 2022/2/28 18:47
* @Desc:
 */

type BlogSection struct {
	Id             int       `db:"id" json:"id"`
	Title          string    `db:"title" json:"title"`
}

type Blog struct {
	BlogSection
	Content        string    `db:"content" json:"content"`
	FirstPicture   string    `db:"first_picture" json:"firstPicture"`
	Flag           string    `db:"flag" json:"flag"`
	Views          int       `db:"views" json:"views"`
	Appreciation   bool      `db:"appreciation" json:"appreciation"`
	ShareStatement bool      `db:"share_statement" json:"shareStatement"`
	Commentabled   bool      `db:"commentabled" json:"commentabled"`
	Published      bool      `db:"published" json:"published"`
	Recommend      bool      `db:"recommend" json:"recommend"`
	CreateTime     time.Time `db:"create_time" json:"createTime"`
	UpdateTime     time.Time `db:"update_time" json:"updateTime"`
	Description    string    `db:"description" json:"description"`
	TypeId         int       `db:"type_id" json:"typeId"`
	UserId         int       `db:"user_id" json:"userId"`
}

type BlogUserType struct {
	Blog
	Nickname string `db:"nickname" json:"nickname"`
	Typename string `db:"typename" json:"typename"`
}

type DetailedBlog struct {
	Blog
	Nickname string `db:"nickname" json:"nickname"`
	Typename string `db:"typename" json:"typename"`
	Views    int    `db:"views" json:"views"`
}

type FullBlog struct {
	BlogUserType
	TagIds []int `json:"tagIds"`
}
