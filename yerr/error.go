package yerr

import (
	"github.com/link-yundi/ytools"
	"github.com/link-yundi/ytools/ylog"
	"github.com/pkg/errors"
)

/**
------------------------------------------------
Created on 2022-11-18 11:03
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

var (
	errChan = make(chan error)
	errList []error
)

func New(err error) error {
	return errors.Wrap(err, ytools.ErrStackTraceSplit)
}

// 错误归集器
func Put(errs ...error) {
	for _, err := range errs {
		if err != nil {
			errChan <- err
		}
	}
}

// 打印error 并且关闭
func HandleFunc(handler func(error)) {
	for _, err := range errList {
		handler(err)
	}
	close(errChan)
}

func listen() {
	for err := range errChan {
		if err != nil {
			errList = append(errList, err)
		}
	}
}

func init() {
	ylog.Info("启动yerror错误收集")
	go listen()
}
