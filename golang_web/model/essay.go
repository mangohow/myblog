package model

import "time"

type Essay struct {
	Id         int       `json:"id" db:"id"`
	CreateTime time.Time `json:"createTime" db:"create_time"`
	Content    string    `json:"content" db:"content"`
}
