package ydir

import (
	"os"
)

/**
------------------------------------------------
Created on 2022-11-07 11:28
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// MkDirs 创建目录
func MkDirs(dirPath string) error {
	if !Has(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		return err
	}
	return nil
}

// Has 判断文件夹是否存在
func Has(dirPath string) bool {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
