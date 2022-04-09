package dao

import (
	"blog_web/utils"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

/*
* @Author: mgh
* @Date: 2022/2/23 21:32
* @Desc:
 */

var sqldb *sqlx.DB
var Sqldb *sqlx.DB

//var conf = &utils.MysqlConf{
//	DataSourceName: "root:mgh99999@(192.168.226.128:3306)/blog?charset=utf8mb4&parseTime=true&loc=Local",
//	MaxOpenConns:   20,
//	MaxIdleConns:   10,
//}

var conf = &utils.GlobalServerConf.Mysql

func init() {
	fmt.Println("mysql conf:", conf)
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	var err error
	sqldb, err = sqlx.ConnectContext(timeoutCtx, "mysql", conf.DataSourceName)
	if err != nil {
		panic(err)
	}
	Sqldb = sqldb
	sqldb.SetMaxOpenConns(int(conf.MaxOpenConns))
	sqldb.SetMaxIdleConns(int(conf.MaxIdleConns))
	fmt.Println("Connect mysql success")
}

func CloseMysqlConn() {
	sqldb.Close()
}
