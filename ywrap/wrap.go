package ywrap

import (
	"reflect"
	"time"
	"ytools/ylog"
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

func WithApply(wrapper interface{}, slice interface{}) {
	if reflect.ValueOf(slice).Kind() != reflect.Slice {
		return
	}
	value := reflect.ValueOf(slice)
	length := value.Len()
	for i := 0; i < length; i++ {
		val := value.Index(i)
		res := reflect.ValueOf(wrapper).Call([]reflect.Value{val})[0]
		val.Set(res)
	}
}
