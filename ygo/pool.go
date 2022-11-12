package ygo

import (
	"runtime"
	"sync"
)

/**
------------------------------------------------
Created on 2022-11-12 19:06
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

var (
	workingChan chan struct{}
	wg          *sync.WaitGroup
)

func init() {
	// 设置可调度的最大核数
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)
	workingChan = make(chan struct{}, cpuNum)
	wg = &sync.WaitGroup{}
}

type Fish struct {
	Arg interface{}
	fn  func(interface{})
}

func NewFish(fn func(interface{})) *Fish {
	return &Fish{
		fn: fn,
	}
}

// 提交协程
func Submit(fns ...*Fish) {
	for _, fn_ := range fns {
		workingChan <- struct{}{}
		wg.Add(1)
		go func(fn *Fish) {
			fn.fn(fn.Arg)
			wg.Done()
			<-workingChan
		}(fn_)
	}
}

func Wait() {
	wg.Wait()
}
