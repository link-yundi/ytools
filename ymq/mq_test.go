package ymq

import (
	"github.com/link-yundi/ytools/ygo"
	"github.com/link-yundi/ytools/ylog"
	"testing"
	"time"
)

/**
------------------------------------------------
Created on 2022-11-12 17:16
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

func TestMq(t *testing.T) {
	defer ygo.Wait()
	topic1 := NewTopic("Topic1", "测试1")
	topic2 := NewTopic("Topic2", "测试2")
	consumer1 := func(msg interface{}) {
		ylog.Info("I'm Consumer1", msg)
	}
	consumer2 := func(msg interface{}) {
		ylog.Info("I'm Consumer2", msg)
		time.Sleep(3 * time.Second)
	}
	Subscribe(topic1, consumer1, consumer2)
	Subscribe(topic2, consumer2)
	Produce(topic1, 1, 2)
	Produce(topic2, "Hello, Fish")
}
