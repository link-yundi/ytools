package yerror

import "github.com/link-yundi/ytools/ylog"

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

// 错误归集器
func Collect(errs ...error) {
	for _, err := range errs {
		if err != nil {
			errChan <- err
		}
	}
}

// 打印error 并且关闭
func Log() {
	for _, err := range errList {
		ylog.Error(err)
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
