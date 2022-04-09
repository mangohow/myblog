package model

import "time"

/*
* @Author: mgh
* @Date: 2022/2/28 18:52
* @Desc:
 */

type User struct {
	Id         int       `db:"id" json:"id"`
	Nickname   string    `db:"nickname" json:"nickname"`
	Username   string    `db:"username" json:"username"`
	Password   string    `db:"password" json:"password"`
	Email      string    `db:"email" json:"email"`
	Avatar     string    `db:"avatar" json:"avatar"`
	Type       int       `db:"type" json:"type"`
	Github     string    `db:"github" json:"github"`
	Csdn       string    `db:"csdn" json:"csdn"`
	CreateTime time.Time `db:"create_time" json:"createTime"`
	UpdateTime time.Time `db:"update_time" json:"updateTime"`
}
