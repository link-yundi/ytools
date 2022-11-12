package ymq

import "sync/atomic"

/**
------------------------------------------------
Created on 2022-11-12 16:27
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

type topic struct {
	Id          int64
	String      string
	Description string
}

func NewTopic(str string, description string) *topic {
	id := atomic.AddInt64(&globalTopicId, 1)
	return &topic{Id: id, String: str, Description: description}
}
