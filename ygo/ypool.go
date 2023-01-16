package ygo

import (
	"reflect"
	"runtime"
	"sync"
	"ytools/ylog"
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
	maxGoOnce   sync.Once
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

// 设置最大协程数
func MaxGo(num int) {
	maxGoOnce.Do(func() {
		ylog.Info("ygo 设置最大并发数为:", num)
		workingChan = make(chan struct{}, num)
	})
}

func Wait() {
	wg.Wait()
}

func Go(fn interface{}, args ...interface{}) {
	fv := reflect.ValueOf(fn)
	switch fv.Kind() {
	case reflect.Func:
		ins := make([]reflect.Value, 0)
		for _, arg := range args {
			ins = append(ins, reflect.ValueOf(arg))
		}
		workingChan <- struct{}{}
		wg.Add(1)
		go func() {
			fv.Call(ins)
			wg.Done()
			<-workingChan
		}()
	}
}
