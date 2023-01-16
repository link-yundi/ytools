package ymath

import "math"

/**
------------------------------------------------
Created on 2022-11-07 16:18
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// 返回 val 根据指定精度 Precision (十进制小数点后数字的数目) 进行四舍五入的结果
func Round(val float64, decimal int) float64 {
	p := math.Pow10(decimal)
	return math.Floor(val*p+0.5) / p
}

// 向下截取精度
func RoundFloor(val float64, decimal int) float64 {
	p := math.Pow10(decimal)
	return math.Floor(val*p) / p
}

// 向上截取精度
func RoundCeil(val float64, decimal int) float64 {
	p := math.Pow10(decimal)
	return math.Ceil(val*p) / p
}
