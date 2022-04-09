package model

import "time"

// 留言

type Message struct {
	Id          int       `json:"id" db:"id"`
	Nickname    string    `json:"nickname" db:"nickname"`
	Email       string    `json:"email" db:"email"`
	Url         string    `json:"url" db:"url"`
	Content     string    `json:"content" db:"content"`
	Avatar      int       `json:"avatar" db:"avatar"`
	CreateTime  time.Time `json:"createTime" db:"create_time"`
	ParentId    int       `json:"parentId" db:"parent_id"`
	TopParentId int       `json:"topParentId" db:"top_parent_id"`
	Status      int       `json:"status" db:"status"`
	Ip          string    `json:"ip" db:"ip"`
}

