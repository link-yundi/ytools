package yerr

import (
	"github.com/pkg/errors"
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
