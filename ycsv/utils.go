package ycsv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/jszwec/csvutil"
	"os"
	"path"
	"sync"
	"ytools/yerr"
	"ytools/ypath"
)

/**
------------------------------------------------
Created on 2022-11-07 11:21
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

var (
	lock = &sync.Mutex{}
)

type Row []interface{}

// 写入csv
func Write(data interface{}, filePath string) error {
	lock.Lock()
	dirPath, _ := path.Split(filePath)
	if !ypath.Has(dirPath) {
		err := ypath.MkDirs(dirPath)
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return yerr.New(err.Error())
	}
	writer := csv.NewWriter(file)
	defer func() {
		writer.Flush()
		file.Close()
		lock.Unlock()
	}()
	encoder := csvutil.NewEncoder(writer)
	err = encoder.Encode(data)
	if err != nil {
		return yerr.New(err.Error())
	}
	return nil
}

// 追加写入
func WriteAppend(data interface{}, filePath string) error {
	lock.Lock()
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return yerr.New(err.Error())
	}
	writer := csv.NewWriter(file)
	defer func() {
		writer.Flush()
		file.Close()
		lock.Unlock()
	}()
	encoder := csvutil.NewEncoder(writer)
	encoder.AutoHeader = false
	err = encoder.Encode(data)
	if err != nil {
		return yerr.New(err.Error())
	}
	return nil
}

// 逐行写入
func WriteRows(data []*Row, filePath string) error {
	lock.Lock()
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return yerr.New(err.Error())
	}
	writer := csv.NewWriter(file)
	defer func() {
		writer.Flush()
		file.Close()
		lock.Unlock()
	}()
	strList := make([][]string, 0)
	for _, row := range data {
		tmp := make([]string, 0)
		for _, elem := range *row {
			tmp = append(tmp, fmt.Sprintf("%v", elem))
		}
		strList = append(strList, tmp)
	}
	if len(strList) > 0 {
		err = writer.WriteAll(strList)
		if err != nil {
			return yerr.New(err.Error())
		}
	}
	return nil
}

// Read 读取csv
func Read(dest interface{}, filePath string) error {
	// dest: 接收的对象
	byteData, err := os.ReadFile(filePath)
	if err != nil {
		return yerr.New(err.Error())
	}
	err = gocsv.UnmarshalBytes(byteData, dest)
	if err != nil {
		return yerr.New(err.Error())
	}
	return nil
}

// 将数据解析成 []byte
func Decode(v interface{}, tags ...string) ([]byte, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	enc := csvutil.NewEncoder(writer)
	for _, tag := range tags {
		enc.Tag += tag
	}
	err := enc.Encode(v)
	if err != nil {
		return nil, yerr.New(err.Error())
	}
	writer.Flush()
	return buf.Bytes(), nil
}

// 将 []byte 解析
func Encode(dest interface{}, byteData []byte) error {
	//err := csvutil.Unmarshal(byteData, dest)
	err := gocsv.UnmarshalBytes(byteData, dest)
	if err != nil {
		return yerr.New(err.Error())
	}
	return nil
}
