package ycsv

import (
	"strconv"
	"sync"
	"testing"
	"time"
	"ytools/ylog"
	"ytools/ypath"
	"ytools/ytime"
)

/**
------------------------------------------------
Created on 2022-11-08 13:37
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

type Student struct {
	Datetime string `csv:"datetime" db:"Datetime"`
	Name     string `csv:"name" db:"Name"`
	Age      int    `csv:"age" db:"Age"`
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
	err := Write(students, "students.csv")
	if err != nil {
		ylog.Error(err)
	}
}

func TestHasFile(t *testing.T) {
	res := ypath.Has("/Users/zhangyundi/go/src/ytools/ycsv/students.csv")
	ylog.Info(res)
}

func TestWriteAppend(t *testing.T) {
	students := make([]Student, 0)
	for i := 0; i < 10; i++ {
		stu := &Student{}
		stu.Name = "zhangyundi" + strconv.FormatInt(int64(i), 10)
		stu.Age = i
		stu.Datetime = ytime.DatetimeMs(time.Now())
		students = append(students, *stu)
	}
	err := WriteAppend(&students, "students.csv")
	if err != nil {
		ylog.Error(err)
	}
}

func TestWriteRows(t *testing.T) {
	dataList := make([]*Row, 0)
	// 表头
	dataList = append(dataList, &Row{"Name", "Age", "Datetime"})
	for i := 0; i < 10; i++ {
		row := &Row{"zhangyundi", i, time.Now()}
		dataList = append(dataList, row)
	}
	err := WriteRows(dataList, "students.csv")
	if err != nil {
		ylog.Error(err)
	}
}

// 解码测试
func TestDecode(t *testing.T) {
	students := make([]Student, 0)
	for i := 0; i < 10; i++ {
		stu := &Student{}
		stu.Name = "zhangyundi" + strconv.FormatInt(int64(i), 10)
		stu.Age = i
		stu.Datetime = ytime.DatetimeMs(time.Now())
		students = append(students, *stu)
	}
	var (
		byteData []byte
		err      error
	)
	byteData, err = Decode(students, "db")
	if err != nil {
		ylog.Fatal(err)
	}
	ylog.Info(string(byteData))
}
