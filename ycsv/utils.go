package ycsv

import (
	"encoding/csv"
	"github.com/go-gota/gota/dataframe"
	"github.com/jszwec/csvutil"
	"os"
)

/**
------------------------------------------------
Created on 2022-11-07 11:21
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

//保存达到文件
func WriteCsv(data interface{}, filePath string) error {
	df := dataframe.LoadStructs(data)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	defer file.Close()
	if err != nil {
		return err
	}
	err = df.WriteCSV(file)
	if err != nil {
		return err
	}
	return nil
}

// 根据结构体的tag保存
func WriteCsvWithTag(data interface{}, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	encoder := csvutil.NewEncoder(writer)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return err
}

// ReadCsv 读取csv
func ReadCsv(filePath string, dest interface{}) error {
	// dest: 接收的对象
	byteData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = csvutil.Unmarshal(byteData, dest)
	if err != nil {
		return err
	}
	return nil
}
