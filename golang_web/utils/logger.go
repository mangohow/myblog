package utils

import (
	"blog_web/utils/logger"
	"github.com/spf13/viper"
)

/*
* @Author: mgh
* @Date: 2022/3/5 13:13
* @Desc:
 */

var logg *logger.Logger

func Logger() *logger.Logger {
	return logg
}

func CreateLogger() {
	if viper.GetBool("log.toFile") {
		logg = logger.NewFileLogger(viper.GetString("log.filepath"),
			viper.GetString("log.filename"),
			viper.GetString("log.level"),
		)
	} else {
		logg = logger.NewLogger(viper.GetString("log.level"))
	}
}
