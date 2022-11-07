package ylog

import "github.com/link-yundi/ytools/ylog"

/**
------------------------------------------------
Created on 2022-11-07 16:12
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// 设置log级别
func ExampleSetLogLevel() {
	ylog.SetLogLevel(ylog.LevelWarn)
	ylog.Info("Info msg")
	ylog.Warn("Warn msg")
	// Output:
	// Warn msg
}
