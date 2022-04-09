package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

/*
* @Author: mgh
* @Date: 2022/3/8 17:07
* @Desc:
 */

type LogLevel uint8

const (
	DEBUG = iota + 1
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	Level       LogLevel
	LevelStr    string
	FilePath    string
	FileName    string
	MaxFileSize uint32
	Out         io.Writer
	Formatter   LogFormatter
}

type logInfo struct {
	level     LogLevel
	strLevel  string
	message   string
	timeStamp string
	fileName  string
	funcName  string
	lineNo    int
}

type LogFormatter func(info *logInfo) string

var DefaultFormatter = func(info *logInfo) string {
	return fmt.Sprintf("[%s][%s][%s:%s:%d] %s\n",
		info.timeStamp,
		info.strLevel,
		info.fileName,
		info.funcName,
		info.lineNo,
		info.message,
	)
}

func NewLogger(level string) *Logger {
	return &Logger{
		Level:     parseLogLevel(level),
		LevelStr:  level,
		Out:       os.Stdout,
		Formatter: DefaultFormatter,
	}
}

func NewFileLogger(filepath, filename, level string) *Logger {
	if !strings.HasSuffix(filename, ".log") {
		filename += ".log"
	}
	out, err := os.OpenFile(path.Join(filepath, filename), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return &Logger{
		Level:     parseLogLevel(level),
		LevelStr:  level,
		FilePath:  filepath,
		FileName:  filename,
		Out:       out,
		Formatter: DefaultFormatter,
	}
}

func (l *Logger) writeLog(level LogLevel, strLevel, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	strNow := time.Now().Format("2006-01-02 15:04:05")
	fileName, funcName, lineNo := getInfo(3)
	fmt.Fprint(l.Out, l.Formatter(&logInfo{
		level:     level,
		strLevel:  strLevel,
		message:   message,
		timeStamp: strNow,
		fileName:  fileName,
		funcName:  funcName,
		lineNo:    lineNo,
	}))
}

func (l *Logger) Debug(format string, args ...interface{}) {
	if l.enable(DEBUG) {
		l.writeLog(DEBUG, "DEBUG", format, args...)
	}
}

func (l *Logger) Info(format string, args ...interface{}) {
	if l.enable(INFO) {
		l.writeLog(INFO, "INFO", format, args...)
	}
}

func (l *Logger) Warning(format string, args ...interface{}) {
	if l.enable(WARNING) {
		l.writeLog(WARNING, "WARN", format, args...)
	}
}

func (l *Logger) Error(format string, args ...interface{}) {
	if l.enable(ERROR) {
		l.writeLog(ERROR, "ERROR", format, args...)
	}
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	if l.enable(FATAL) {
		l.writeLog(FATAL, "FATAL", format, args...)
	}
}

func (l *Logger) enable(level LogLevel) bool {
	return level >= l.Level
}

// get func name, file name and line number when print the log
func getInfo(skip int) (fileName, funcName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}

	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]

	return
}

// convert string to LogLevel
func parseLogLevel(level string) LogLevel {

	level = strings.ToLower(level)
	switch level {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warning":
		return WARNING
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		panic("Invalid log level")
	}

}
