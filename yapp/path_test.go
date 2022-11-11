package yapp

import (
	"github.com/link-yundi/ytools/ylog"
	"testing"
)

/**
------------------------------------------------
Created on 2022-11-11 14:50
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

func TestInitApp(t *testing.T) {
	err := InitApp("/Users/zhangyundi", "testapp")
	if err != nil {
		ylog.Error(err)
	}
	Info()
}
