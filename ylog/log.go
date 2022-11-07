package ylog

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	log_ "log"
	"os"
	"strings"
)

/**
------------------------------------------------
Created on 2022-11-07 12:16
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

type logLevel int

const (
	LevelTrace logLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	logFormat = log_.Ldate | log_.Ltime | log_.Lshortfile
)

var (
	l = &logger{
		Logger: log_.New(os.Stdout, "", logFormat),
	}
	levelMap = map[string]logLevel{
		"trace": LevelTrace,
		"debug": LevelDebug,
		"info":  LevelInfo,
		"warn":  LevelWarn,
		"error": LevelError,
	}
)

type logger struct {
	level logLevel
	*log_.Logger
}

func SetLogFlags(flag int) {
	l.SetFlags(flag)
}

func SetLogLevel(logLevel logLevel) {
	l.level = logLevel
}

// 通过string设置级别，trace > debug > info > warn > error: 大小写均可
func SetLogLevelFromStr(level string) {
	level = strings.ToLower(level)
	if loglevel, ok := levelMap[level]; ok {
		l.level = loglevel
	}
}

// 设置log文件路径
func SetLogFile(logPath string) {
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		Error(err)
	}
	writers := []io.Writer{
		file, os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	l.Logger = log_.New(fileAndStdoutWriter, "", logFormat)
}

func canInfo(level logLevel) bool {
	return l.level <= level
}

func Trace(v ...interface{}) {
	if !canInfo(LevelTrace) {
		return
	}
	l.Output(2, fmt.Sprintln("[TRACE]", v))
}

func Tracef(format string, v ...interface{}) {
	if !canInfo(LevelTrace) {
		return
	}
	format = "[TRACE] [" + format + "]"
	l.Output(2, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	if !canInfo(LevelDebug) {
		return
	}
	l.Output(2, fmt.Sprintln("[DEBUG]", v))
}

func Debugf(format string, v ...interface{}) {
	if !canInfo(LevelDebug) {
		return
	}
	format = "[DEBUG] [" + format + "]"
	l.Output(2, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	if !canInfo(LevelWarn) {
		return
	}
	l.Output(2, fmt.Sprintln("[WARN]", v))
}

func Warnf(format string, v ...interface{}) {
	if !canInfo(LevelWarn) {
		return
	}
	format = "[WARN] [" + format + "]"
	l.Output(2, fmt.Sprintf(format, v...))
}

func Error(err error) {
	if !canInfo(LevelError) {
		return
	}
	l.Output(2, fmt.Sprintf("[ERROR] original err:%T\n", errors.Cause(err)))
	l.Output(2, fmt.Sprintf("[ERROR] error msg: %v\n", errors.Cause(err)))
	l.Output(2, fmt.Sprintf("[ERROR] stack trace: \n %+v\n", err))
}

func Info(v ...interface{}) {
	if !canInfo(LevelInfo) {
		return
	}
	l.Output(2, fmt.Sprintln("[INFO]", v))
}

func Infof(format string, v ...interface{}) {
	if !canInfo(LevelInfo) {
		return
	}
	format = "[INFO] [" + format + "]"
	l.Output(2, fmt.Sprintf(format, v...))
}

func Fatal(v ...interface{}) {
	l.Output(2, fmt.Sprintln("[FATAL]", v))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	format = "[FATAL] [" + format + "]"
	l.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}
