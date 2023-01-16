package ypath

import (
	"testing"
	"ytools/ylog"
)

/**
------------------------------------------------
Created on 2022-11-07 13:58
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

var dirPath = "./testdir"

func TestHas(t *testing.T) {
	ylog.Info(Has(dirPath))
}

func TestMkDirs(t *testing.T) {
	MkDirs(dirPath)
}
