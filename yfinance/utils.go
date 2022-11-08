package yfinance

import "github.com/link-yundi/ytools/ymath"

/**
------------------------------------------------
Created on 2022-11-07 19:46
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// 浮点数偏移：比如计算 价格 偏移 一定后的价格，精度和场景有关，比如a股,精度为2: 0.01
func FloatOffset(p float64, decimal int, offset ...float64) float64 {
	res := p
	for _, o := range offset {
		res = ymath.Round(res+o, decimal)
	}
	return res
}
