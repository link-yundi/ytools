package ywrap

import (
	"fmt"
	"github.com/link-yundi/ytools/ylog"
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
	panicFunc := func() {
		defer ylog.Panic()
		a := make([]int, 0)
		fmt.Println(a[0])

	}
	//WithRecover(panicFunc)
	panicFunc()
	fmt.Println("go on 恢复运行")
}

func TestWithApply(t *testing.T) {
	a := []int{1, 3}
	WithApply(func(x int) int { return x + 1 }, a)
	fmt.Println(a)
}
