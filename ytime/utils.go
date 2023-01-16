package ytime

import (
	"fmt"
	"strconv"
	"time"
	"ytools/yerr"
)

/**
------------------------------------------------
Created on 2022-11-07 16:31
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

const (
	DateLayout       string = "20060102"
	DatetimeLayout   string = "2006-01-02 15:04:05"
	DatetimeMsLayout string = "2006-01-02 15:04:05.000"
	DateLayout_      string = "2006-01-02"
	TimeLayout       string = "15:04:05"
	TimeMsLayout     string = "15:04:05.000"
)

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
func TimeOffsetDays(t time.Time, days int) time.Time {
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
func TimeOffsetSecs(t time.Time, secs int) time.Time {
	res := time.Unix(t.Unix()+int64(secs), 0).UTC()
	return res
}

// 时间 加减 毫秒
func TimeOffsetMilli(t time.Time, milli int) time.Time {
	msec := UnixMilli(t) + int64(milli)
	return time.Unix(msec/1e3, (msec%1e3)*1e6).UTC()
}

func UnixMilli(t time.Time) int64 {
	return int64(t.UnixNano()) / 1e6
}

// 获取时间列表
func TimeList(startT, endT time.Time, step int, offsetFunc func(t time.Time, offset int) time.Time) []time.Time {
	res := make([]time.Time, 0)
	nextT := startT
	for nextT.Before(endT) {
		res = append(res, nextT)
		nextT = offsetFunc(nextT, step)
	}
	res = append(res, endT)
	return res
}

// 时间解析
func Parse(layout string, t string) (time.Time, error) {
	res, err := time.Parse(layout, t)
	if err != nil {
		return time.Time{}, yerr.New(err.Error())
	}
	return res, nil
}

// 规范输出
func Output(secs float64) (str string) {
	daySecs := 3600 * 24
	d := int(secs) / daySecs
	h := int(secs) % daySecs / 3600
	m := int(secs) % 3600 / 60
	s := int(secs) % 60
	ms := int(secs*1000) % 1000
	if d > 0 {
		str += strconv.Itoa(d) + "[d]"
	}
	if h > 0 {
		str += strconv.Itoa(h) + "[h]"
	}
	if m > 0 {
		str += strconv.Itoa(m) + "[m]"
	}
	if s > 0 {
		str += strconv.Itoa(s) + "[s]"
	}
	if ms > 0 {
		str += strconv.Itoa(ms) + "[ms]"
	}
	return
}

// duration
func Duration(layout, startTime, endTime string) (float64, error) {
	var (
		startT, endT time.Time
		err          error
	)
	startT, err = Parse(layout, startTime)
	if err != nil {
		return 0, err
	}
	endT, err = Parse(layout, endTime)
	if err != nil {
		return 0, err
	}
	duration := float64(UnixMilli(endT)-UnixMilli(startT)) / 1000
	return duration, nil
}
