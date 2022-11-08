package ycsv

import (
	"github.com/link-yundi/ytools/ylog"
	"github.com/link-yundi/ytools/ypath"
	"github.com/link-yundi/ytools/ytime"
	"strconv"
	"sync"
	"testing"
	"time"
)

/**
------------------------------------------------
Created on 2022-11-08 13:37
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

type Student struct {
	Datetime string `csv:"datetime"`
	Name     string `csv:"name"`
	Age      int    `csv:"age"`
}

func TestWriteCsv(t *testing.T) {
	stuPool := &sync.Pool{
		New: func() interface{} {
			return &Student{}
		},
	}
	students := make([]Student, 0)
	for i := 0; i < 10; i++ {
		stu := stuPool.Get().(*Student)
		stu.Name = "zhangyundi" + strconv.FormatInt(int64(i), 10)
		stu.Age = i
		stu.Datetime = ytime.DatetimeMs(time.Now())
		students = append(students, *stu)
		stuPool.Put(stu)
	}
	err := WriteCsvWithTag(students, "students.csv")
	if err != nil {
		ylog.Error(err)
	}
}

func TestHasFile(t *testing.T) {
	res := ypath.Has("/Users/zhangyundi/go/src/github.com/link-yundi/ytools/ycsv/students.csv")
	ylog.Info(res)
}
