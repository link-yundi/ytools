package yerr

import (
	"fmt"
	"github.com/pkg/errors"
	"ytools/ylog"
)

/**
------------------------------------------------
Created on 2022-11-18 11:03
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

var errHandler func(error)

func New(message string) error {
	return errors.New(message)
}

// 错误归集器
func Put(errs ...error) {
	for _, err := range errs {
		if err == nil {
			continue
		}
		if errHandler != nil {
			errHandler(err)
		}
	}
}

// 打印error 并且关闭
func HandleFunc(handleFunc func(error)) {
	errHandler = handleFunc
}

// ========================== 一些预定义的错误 ==========================
var (
	LocalDataFileNotExist = New("本地数据不存在")
)

// 错误统一处理
func ErrHandler(err error) {
	switch {
	case errors.Is(err, LocalDataFileNotExist):
		ylog.Output(4, fmt.Sprintf("[ERROR] Msg: %v\n", err))
	default:
		ylog.Output(4, fmt.Sprintf("[ERROR] %T\n %+v\n", errors.Cause(err), err))
	}
}
