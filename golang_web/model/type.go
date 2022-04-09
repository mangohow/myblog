package model


/*
* @Author: mgh
* @Date: 2022/2/28 18:55
* @Desc:
 */

type TheType struct {
	Id    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Count int    `db:"count" json:"count"`
}
