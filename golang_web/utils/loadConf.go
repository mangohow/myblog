package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

/*
* @Author: mgh
* @Date: 2022/3/1 19:21
* @Desc:
 */

func init() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("an error occurs when load conf file:%v", err))
	}

	CreateLogger()
}
