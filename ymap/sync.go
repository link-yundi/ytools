package ymap

import (
	"sync"
)

/**
------------------------------------------------
Created on 2022-12-16 10:49
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// ========================== 基于sync.Map 嵌套 map 的并发安全 ==========================
type Map struct {
	*sync.Map // map[key]*childMap
}

func NewMap() *Map {
	return &Map{
		Map: &sync.Map{},
	}
}

func (m *Map) has(key interface{}) bool {
	if _, ok := m.Map.Load(key); !ok {
		return false
	}
	return true
}

func (m *Map) Has(keys ...interface{}) bool {
	iter := m
	for _, key := range keys {
		if key == nil {
			continue
		}
		res := iter.load(key)
		if res == nil {
			return false
		}
		iter = res.(*Map)
	}
	return true
}

// 多个key，制造多个嵌套的 sync.Map
func (m *Map) Store(v interface{}, keys ...interface{}) {
	var key interface{}
	iter := m
	for _, k := range keys {
		if k == nil {
			continue
		}
		iter = iter.newIfNotExist(k)
		key = k
	}
	if key != nil {
		iter.Map.Store(key, v)
	}
}

func (m *Map) Append(v interface{}, keys ...interface{}) {
	if !m.Has(keys...) {
		m.Store(&[]interface{}{}, keys...)
	}
	target := *(m.Load(keys...).(*[]interface{}))
	target = append(target, v)
	m.Store(&target, keys...)
}

func (m *Map) Load(keys ...interface{}) interface{} {
	var key interface{}
	iter := m
	for _, k := range keys {
		if k == nil {
			continue
		}
		res := iter.load(k)
		if res == nil {
			return nil
		}
		iter = res.(*Map)
		key = k
	}
	if key != nil {
		if final, ok := iter.Map.Load(key); !ok {
			return iter
		} else {
			return final
		}
	}
	return nil
}

func (m *Map) Range() map[interface{}]interface{} {
	res := make(map[interface{}]interface{}, 0)
	m.Map.Range(func(key, value interface{}) bool {
		switch value.(type) {
		case *Map:
			res[key] = value.(*Map).load(key)
		default:
			res[key] = value
		}
		return true
	})
	return res
}

// 如果没有指定的key，则返回nil
func (m *Map) load(key interface{}) interface{} {
	if !m.has(key) {
		return nil
	}
	res, _ := m.Map.Load(key)
	return res
}

// 嵌套 map
func (m *Map) newMap(k interface{}) *Map {
	res := &Map{
		Map: &sync.Map{},
	}
	m.Map.Store(k, res)
	return res
}

// 如果不存在则嵌套新的 map
func (m *Map) newIfNotExist(k interface{}) *Map {
	if !m.has(k) {
		return m.newMap(k)
	}
	return m.load(k).(*Map)
}

// 获取key的数量
func (m *Map) Len() int {
	var length int
	m.Map.Range(func(key, value any) bool {
		length++
		return true
	})
	return length
}
