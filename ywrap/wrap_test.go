package ywrap

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

/**
------------------------------------------------
Created on 2022-11-08 00:08
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

func add(a, b int) int {
	return a + b
}

func sleep() {
	time.Sleep(3 * time.Second)
}

func TestWrapTime(t *testing.T) {
	res := WithTime(add, 1, 2)
	fmt.Println(res[0].(int))
}

func TestNilRet(t *testing.T) {
	WithTime(sleep)
}

func TestPanic(t *testing.T) {
	err := errors.New("panic")
	WithRecover(func() { panic(err) })
}
