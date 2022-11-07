package ytime

import (
	"github.com/link-yundi/ytools/ylog"
	"testing"
	"time"
)

/**
------------------------------------------------
Created on 2022-11-07 18:59
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

func TestTimeList(t *testing.T) {
	startT := time.Now()
	endT := SecsOffset(startT, 30)
	tList := TimeList(startT, endT, 4, SecsOffset)
	for _, d := range tList {
		ylog.Info(Datetime(d))
	}
}
