package ymq

import (
	"github.com/link-yundi/ytools/ygo"
	"sync"
)

/**
------------------------------------------------
Created on 2022-11-07 23:15
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

var (
	consumerMap         = &sync.Map{}
	globalTopicId int64 = 0
)

func Produce(topic *topic, msgs ...interface{}) {
	if v, ok := consumerMap.Load(topic.Id); ok {
		fishList := v.(*[]*ygo.Fish)
		for _, msg := range msgs {
			for _, fish := range *fishList {
				fish.Arg = msg
				ygo.Submit(fish)
			}
		}
	}
}

func Subscribe(topic *topic, consumers ...func(interface{})) {
	if _, ok := consumerMap.Load(topic.Id); !ok {
		consumerMap.Store(topic.Id, new([]*ygo.Fish))
	}
	v, _ := consumerMap.Load(topic.Id)
	consumerList := v.(*[]*ygo.Fish)
	for _, consumer := range consumers {
		fish := ygo.NewFish(consumer)
		*consumerList = append(*consumerList, fish)
	}
}
