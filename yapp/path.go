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
	AppName string
	HomeDir string // 用户家目录，绝对路径，由调用者设置
	AppDir  string // app 目录 HomeDir/applications/app-name
	LogDir  string // 日志文件目录 HomeDir/var/app-name/log
	ConfDir string // 配置文件目录路径 HomeDir/etc/app-name/conf
)

// app 初始化
func InitApp(userHomePath, appname string) error {
	var err error
	AppName = appname
	HomeDir = userHomePath
	AppDir = path.Join(HomeDir, "applications", AppName)
	LogDir = path.Join(HomeDir, "var", AppName, "log")
	ConfDir = path.Join(HomeDir, "etc", AppName, "conf")
	for _, dir := range []string{HomeDir, AppDir, LogDir, ConfDir} {
		err = ypath.MkDirs(dir)
		if err != nil {
			return err
		}
	}
	Info()
	return nil
}

// 打印app相关目录
func Info() {
	ylog.Info("AppName:", AppName)
	ylog.Info("HomeDir:", HomeDir)
	ylog.Info("AppDir:", AppDir)
	ylog.Info("LogDir:", LogDir)
	ylog.Info("ConfDir:", ConfDir)
}
