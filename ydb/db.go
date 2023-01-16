package ydb

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"reflect"
	"strings"
	"time"
	"ytools/ycsv"
	"ytools/yerr"
	"ytools/ylog"
	"ytools/yreflect"
)

/**
------------------------------------------------
Created on 2022-12-19 13:45
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

type Options struct {
	Addrs        []string
	Database     string
	UserName     string
	Password     string
	MaxIdleConns int
	MaxOpenConns int
	Driver       string
	Tag          string
}

type DB struct {
	*sqlx.DB
}

func NewDb(opt *Options) *DB {
	var db *sql.DB
	switch opt.Driver {
	case "clickhouse":
		db = clickhouse.OpenDB(&clickhouse.Options{
			Addr: opt.Addrs,
			Auth: clickhouse.Auth{
				Database: opt.Database,
				Username: opt.UserName,
				Password: opt.Password,
			},
			DialTimeout: 10 * time.Second,
			Compression: &clickhouse.Compression{
				Method: clickhouse.CompressionLZ4,
			},
			Debug:           false,
			BlockBufferSize: 10,
		})
	case "mysql":
		cfg := &mysql.Config{
			Net:                  "tcp",
			User:                 opt.UserName,
			Passwd:               opt.Password,
			Addr:                 opt.Addrs[0],
			DBName:               opt.Database,
			Timeout:              10 * time.Second,
			AllowNativePasswords: true,
			Collation:            "utf8mb4_general_ci",
		}
		conn, err := mysql.NewConnector(cfg)
		if err != nil {
			ylog.Fatal("mysql 连接失败", err)
		}
		db = sql.OpenDB(conn)
	default:
		ylog.Fatal("未知的 db driver", opt.Driver)
	}

	xdb := sqlx.NewDb(db, opt.Driver)
	if opt.MaxIdleConns > 0 {
		xdb.SetMaxIdleConns(opt.MaxIdleConns)
	}
	if opt.MaxOpenConns > 0 {
		xdb.SetMaxOpenConns(opt.MaxOpenConns)
	}
	if opt.Tag != "" {
		xdb.Mapper = reflectx.NewMapperFunc(opt.Tag, func(s string) string {
			return s
		})
	}
	return &DB{
		DB: xdb,
	}
}

func (db *DB) Get(dest interface{}, query string) error {
	err := db.Select(dest, query)
	if err != nil {
		return yerr.New(err.Error())
	}
	return nil
}

// tb: 如果带有dbName, tb 为 {db}.{tb} 格式
func (db *DB) Put(v interface{}, tb string) error {
	structList := yreflect.WalkSliceVal(reflect.ValueOf(v))
	byteData, err := ycsv.Decode(v, "ch")
	if err != nil {
		return err
	}
	rows := strings.Split(string(byteData), "\n")
	header := rows[0]
	insertStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (", tb, header)
	headerList := strings.Split(header, ",")
	for i, field := range headerList {
		insertStr += ":" + field
		if i < len(headerList)-1 {
			insertStr += ", "
		} else {
			insertStr += ")"
		}
	}
	_, err = db.NamedExec(insertStr, structList)
	if err != nil {
		return yerr.New(err.Error())
	}
	return nil
}

func (db *DB) Close() {
	if db.DB != nil {
		db.DB.Close()
		db.DB = nil
	}
}
