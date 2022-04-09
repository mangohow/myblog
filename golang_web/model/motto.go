package model

import "time"

/*
* @Author: mgh
* @Date: 2022/3/19 11:39
* @Desc:
 */

type Motto struct {
	Id         uint32    `db:"id" json:"id"`
	Ch         string    `db:"ch" json:"ch"`
	En         string    `db:"en" json:"en"`
	CreateTime time.Time `json:"createTime" db:"create_time"`
}
