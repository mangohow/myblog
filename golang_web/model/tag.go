package model

/*
* @Author: mgh
* @Date: 2022/2/28 18:57
* @Desc:
 */

type Tag struct {
	Id    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	count int    `db:"count" json:"count"`
}

type BlogTag struct {
	BlogId int `db:"blog_id"`
	TagId  int `db:"tag_id"`
}

