package examples

import "github.com/link-yundi/ytools/ylog"

/**
------------------------------------------------
Created on 2022-11-07 15:45
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// 设置log
func ExampleSetLogLevel() {
	ylog.SetLogLevel(ylog.LevelWarn)
	ylog.Info("Info msg")
	ylog.Warn("Warn msg")
	// Output:
	// Warn msg
}
