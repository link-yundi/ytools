package yreflect

import (
	"reflect"
)

/**
------------------------------------------------
Created on 2022-12-19 10:01
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// ========================== interface{} -> []interface{} ==========================
func WalkSliceVal(value reflect.Value) []interface{} {
	switch value.Kind() {
	case reflect.Ptr, reflect.Interface:
		return WalkSliceVal(value.Elem())
	case reflect.Slice, reflect.Array:
		res := make([]interface{}, 0)
		for i := 0; i < value.Len(); i++ {
			res = append(res, value.Index(i).Interface())
		}
		return res
	}
	return []interface{}{}
}

func Len(value reflect.Value) int {
	switch value.Kind() {
	case reflect.Ptr, reflect.Interface:
		return Len(value.Elem())
	case reflect.Slice, reflect.Array, reflect.Map, reflect.String:
		return value.Len()
	default:
		return 0
	}
}
