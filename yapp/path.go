package yapp

import (
	"github.com/link-yundi/ytools/ylog"
	"github.com/link-yundi/ytools/ypath"
	"path"
)

/**
------------------------------------------------
Created on 2022-11-11 14:21
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// ========================== y 项目目录相关 ==========================
var (
	AppName      string
	UserHomePath string // 用户家目录，绝对路径，由调用者设置
	AppDir       string // app 目录 UserHomePath/applications/app-name
	LogDir       string // 日志文件目录 UserHomePath/var/app-name/log
	ConfDir      string // 配置文件目录路径 UserHomePath/etc/app-name/conf
)

// app 初始化
func InitApp(userHomePath, appname string) error {
	var err error
	AppName = appname
	UserHomePath = userHomePath
	AppDir = path.Join(UserHomePath, "applications", AppName)
	LogDir = path.Join(UserHomePath, "var", AppName, "log")
	ConfDir = path.Join(UserHomePath, "etc", AppName, "conf")
	for _, dir := range []string{UserHomePath, AppDir, LogDir, ConfDir} {
		err = ypath.MkDirs(dir)
		if err != nil {
			return err
		}
	}
	return nil
}

// 打印app相关目录
func Info() {
	ylog.Info("AppName:", AppName)
	ylog.Info("UserHomePath:", UserHomePath)
	ylog.Info("AppDir:", AppDir)
	ylog.Info("LogDir:", LogDir)
	ylog.Info("ConfDir:", ConfDir)
}
