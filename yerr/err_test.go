package yerr

import (
	"github.com/pkg/errors"
	"testing"
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
	err1 = errors.New("错误1")
	Collect(err1, err2)
}

func fn2() {
	var err3 error
	err3 = errors.New("错误3")
	Collect(err3)
}

func TestYerr(t *testing.T) {
	// 最外层统一打印error
	//go listen()
	defer Log()
	fn1()
	fn2()
}
