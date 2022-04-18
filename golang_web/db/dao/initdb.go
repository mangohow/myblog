package dao

import (
	_ "blog_web/utils"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"time"
)

/*
* @Author: mgh
* @Date: 2022/2/23 21:32
* @Desc:
 */

var sqldb *sqlx.DB
var Sqldb *sqlx.DB

type MysqlConf struct {
	DataSourceName string `json:"data_source_name"`
	MaxOpenConns   uint32 `json:"max_open_conns"`
	MaxIdleConns   uint32 `json:"max_idle_conns"`
}

var conf = &MysqlConf{
	DataSourceName: viper.GetString("mysql.dataSourceName"),
	MaxOpenConns:   viper.GetUint32("mysql.maxOpenConns"),
	MaxIdleConns:   viper.GetUint32("mysql.maxIdleConns"),
}

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
