package ytime

import (
	"github.com/link-yundi/ytools/ylog"
	"testing"
)

/**
------------------------------------------------
Created on 2022-11-07 18:59
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

func TestTimeList(t *testing.T) {
	startDatetime := "2022-11-18 17:41:00.000"
	startDt, err := Parse(DatetimeMsLayout, startDatetime)
	if err != nil {
		ylog.Error(err)
	}
	endT := TimeOffsetMilli(startDt, 500)
	tList := TimeList(startDt, endT, 50, TimeOffsetMilli)
	for _, d := range tList {
		ylog.Info(DatetimeMs(d))
	}
}
