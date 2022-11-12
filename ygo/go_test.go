package ygo

import (
	"github.com/link-yundi/ytools/ylog"
	"testing"
	"time"
)

/**
------------------------------------------------
Created on 2022-11-12 19:25
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// 协程池
func TestGoPool(t *testing.T) {
	defer Wait()
	task1 := func(interface{}) {
		ylog.Info("I'm Task1")
		time.Sleep(2 * time.Second)
		ylog.Info("Task1 Done")
	}
	fish1 := NewFish(task1)
	task2 := func(interface{}) {
		ylog.Info("I'm Task2")
		time.Sleep(5 * time.Second)
		ylog.Info("Task2 Done")
	}
	fish2 := NewFish(task2)
	Submit(fish1, fish2, fish2, fish2, fish1)
}
