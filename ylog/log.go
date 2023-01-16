package ylog

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	log_ "log"
	"os"
	"runtime"
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
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
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
	str := "[TRACE] " + fmt.Sprintln(v...)
	l.Output(2, str)
}

func Tracef(format string, v ...interface{}) {
	if !canInfo(LevelTrace) {
		return
	}
	format = "[TRACE] " + format
	l.Output(2, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	if !canInfo(LevelDebug) {
		return
	}
	str := "[DEBUG] " + fmt.Sprintln(v...)
	l.Output(2, str)
}

func Debugf(format string, v ...interface{}) {
	if !canInfo(LevelDebug) {
		return
	}
	format = "[DEBUG] " + format
	l.Output(2, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	if !canInfo(LevelWarn) {
		return
	}
	str := "[WARN] " + fmt.Sprintln(v...)
	l.Output(2, str)
}

func Warnf(format string, v ...interface{}) {
	if !canInfo(LevelWarn) {
		return
	}
	format = "[WARN] " + format
	l.Output(2, fmt.Sprintf(format, v...))
}

func Error(err error) {
	if !canInfo(LevelError) {
		return
	}
	l.Output(2, fmt.Sprintf("[ERROR] %T\n %+v\n", errors.Cause(err), err))
}

// 只是输出错误信息，不带堆栈
func ErrorMsg(err error) {
	if !canInfo(LevelError) {
		return
	}
	l.Output(2, fmt.Sprintf("[ERROR] Msg: %v\n", err))
}

func Output(calldepth int, s string) {
	l.Output(calldepth, s)
}

func Info(v ...interface{}) {
	if !canInfo(LevelInfo) {
		return
	}
	str := "[INFO] " + fmt.Sprintln(v...)
	l.Output(2, str)
}

func Infof(format string, v ...interface{}) {
	if !canInfo(LevelInfo) {
		return
	}
	format = "[INFO] " + format
	l.Output(2, fmt.Sprintf(format, v...))
}

func InfoStru(v ...interface{}) {
	for _, d := range v {
		infoStru(d)
	}
}

func infoStru(v interface{}) {
	format := "[INFO] %+v"
	l.Output(3, fmt.Sprintf(format, v))
}

func Panic() {
	if err := recover(); err != nil {
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		l.Output(2, fmt.Sprintf("[PANIC] original err:%v", err))
		l.Output(2, fmt.Sprintf("[PANIC] stack trace: \n %s", string(buf[:n])))
	}
}

func Fatal(v ...interface{}) {
	str := "[FATAL] " + fmt.Sprintln(v...)
	l.Output(2, str)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	format = "[FATAL] " + format
	l.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Print(v ...interface{}) {
	str := "[PRINT] " + fmt.Sprintln(v...)
	l.Output(2, str)
}

func Printf(format string, v ...interface{}) {
	format = "[PRINT] " + format
	l.Output(2, fmt.Sprintf(format, v...))
}
