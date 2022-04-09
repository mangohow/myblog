package utils

import "blog_web/utils/logger"

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
	if GlobalServerConf.Log.ToFile {
		logg = logger.NewFileLogger(GlobalServerConf.Log.Filepath,
			GlobalServerConf.Log.Filename,
			GlobalServerConf.Log.Level,
		)
	} else {
		logg = logger.NewLogger(GlobalServerConf.Log.Level)
	}
}
