package utils

import (
	"encoding/json"
	"os"
)

/*
* @Author: mgh
* @Date: 2022/3/1 19:21
* @Desc:
 */

type ServerConf struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	ImagePath    string `json:"image_path"`       // 图片保存路径
	ImageBaseUrl string `json:"image_base_url"`   // 图片访问Base Url
}

type MysqlConf struct {
	DataSourceName string `json:"data_source_name"`
	MaxOpenConns   uint32 `json:"max_open_conns"`
	MaxIdleConns   uint32 `json:"max_idle_conns"`
}

type LogConf struct {
	Filepath string `json:"filepath"`
	Filename string `json:"filename"`
	ToFile   bool   `json:"to_file"`
	Level    string `json:"level"`
}

type Conf struct {
	Server ServerConf `json:"server"`
	Mysql  MysqlConf  `json:"mysql"`
	Log    LogConf    `json:"log"`
}

var GlobalServerConf Conf

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	confPath := wd + "/conf/conf.json"
	bytes, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &GlobalServerConf)
	if err != nil {
		panic(err)
	}

	CreateLogger()
}
