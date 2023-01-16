package yapp

import "ytools/ylog"

/**
------------------------------------------------
Created on 2022-11-11 14:21
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// ========================== y 项目目录相关 ==========================
var (
	HomeDir   string // 用户家目录，绝对路径，由调用者设置
	LogDir    string // 日志文件目录 AppDir/log
	ConfDir   string // 配置文件目录路径 AppDir/conf
	OutputDir string // 产生数据目录路径 AppDir/var
)

// 打印app相关目录
func Info() {
	ylog.Info("HomeDir:", HomeDir)
	ylog.Info("LogDir:", LogDir)
	ylog.Info("ConfDir:", ConfDir)
	ylog.Info("OutputDir:", OutputDir)
}
