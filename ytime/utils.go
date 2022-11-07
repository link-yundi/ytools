package ytime

import (
	"fmt"
	"time"
)

/**
------------------------------------------------
Created on 2022-11-07 16:31
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// yyyymmdd
func Date(t time.Time) string {
	return t.Format(DateLayout)
}

// yyyy-mm-dd hh:mm:ss
func Datetime(t time.Time) string {
	return t.Format(DatetimeLayout)
}

// yyyy-mm-dd hh:mm:ss.000
func DatetimeMs(t time.Time) string {
	return t.Format(DatetimeMsLayout)
}

// hh:mm:ss
func Time(t time.Time) string {
	return t.Format(TimeLayout)
}

// hh:mm:ss.000
func TimeMs(t time.Time) string {
	return t.Format(TimeMsLayout)
}

// 时间 加减 天数
func DaysOffset(t time.Time, days int) time.Time {
	var (
		durations time.Duration
		endDt     time.Time
		hours     string
	)

	if days > 0 {
		hours = fmt.Sprintf("%dh", 24*days)
	} else if days < 0 {
		hours = fmt.Sprintf("-%dh", 24*days)
	}

	durations, _ = time.ParseDuration(hours)
	endDt = t.Add(durations)
	return endDt
}

// 时间 加减 秒
func SecsOffset(t time.Time, secs int) time.Time {
	res := time.Unix(t.Unix()+int64(secs), 0)
	return res
}

// 时间 加减 毫秒
func MilliOffset(t time.Time, milli int) time.Time {
	msec := UnixMilli(t) + int64(milli)
	return time.Unix(msec/1e3, (msec%1e3)*1e6)
}

func UnixMilli(t time.Time) int64 {
	return t.Unix()*1e3 + int64(t.UnixNano())/1e6
}

// 获取时间列表
func TimeList(startT, endT time.Time, offset int, offsetFunc func(t time.Time, offset int) time.Time) []time.Time {
	res := make([]time.Time, 0)
	nextT := startT
	for nextT.Before(endT) {
		res = append(res, nextT)
		nextT = offsetFunc(nextT, offset)
	}
	res = append(res, endT)
	return res
}
