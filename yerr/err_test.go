package yerr

import (
	"github.com/pkg/errors"
	"testing"
	"ytools/ylog"
)

/**
------------------------------------------------
Created on 2022-11-18 11:07
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

func fn1() {
	var err1, err2 error
	err1 = New("错误1")
	Put(err1, err2)
}

func fn2() {
	var err3 error
	err3 = New("错误3")
	err3 = errors.WithMessage(err3, "额外信息")
	Put(err3)
}

func TestYerr(t *testing.T) {
	// 最外层统一打印error
	//go listen()
	HandleFunc(ylog.ErrorMsg)
	fn1()
	fn2()
}
