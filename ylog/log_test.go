package ylog

import (
	"github.com/pkg/errors"
	"testing"
)

/**
------------------------------------------------
Created on 2022-11-07 12:26
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

func TestDebug(t *testing.T) {
	SetLogLevel(LevelDebug)
	msg := "debug"
	Trace(msg)                 // 不会输出
	Debug(msg)                 // 可以输出
	Info(msg)                  // 可以输出
	Warn(msg)                  // 可以输出
	SetLogLevelFromStr("info") // 通过string 设置
	err := errors.New("test err")
	Trace(err)
	Debug(err)
	Info(err)
	Error(err)
}
