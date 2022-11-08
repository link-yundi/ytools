package ywrap

import (
	"github.com/link-yundi/ytools/ylog"
	"reflect"
	"time"
)

/**
------------------------------------------------
Created on 2022-11-07 23:57
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// 耗时装饰器
func WithTime(wrapper interface{}, args ...interface{}) (outs []interface{}) {
	startT := time.Now()
	var ins []reflect.Value
	for _, arg := range args {
		ins = append(ins, reflect.ValueOf(arg))
	}
	resList := reflect.ValueOf(wrapper).Call(ins)
	for _, r := range resList {
		outs = append(outs, r.Interface())
	}
	endT := time.Now()
	duration := endT.Sub(startT).Seconds()
	ylog.Printf("运行时间: %.3f s", duration)
	return outs
}

func WithRecover(wrapper interface{}, args ...interface{}) (outs []interface{}) {
	defer ylog.Panic()
	var ins []reflect.Value
	for _, arg := range args {
		ins = append(ins, reflect.ValueOf(arg))
	}
	resList := reflect.ValueOf(wrapper).Call(ins)
	for _, r := range resList {
		outs = append(outs, r.Interface())
	}
	return outs
}
