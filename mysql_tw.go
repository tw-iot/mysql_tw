package mysql_tw

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var MysqlTw *sql.DB

type MysqlInfo struct {
	Network         string
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	Charset         string
	Other           string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func NewMysqlInfo(host, username, password, database string, port int) MysqlInfo {
	return MysqlInfo{
		Network:         "tcp",
		Host:            host,
		Port:            port,
		Username:        username,
		Password:        password,
		Database:        database,
		Charset:         "utf8",
		Other:           "loc=Asia%2FShanghai&parseTime=true",
		MaxOpenConns:    100,
		MaxIdleConns:    20,
		ConnMaxLifetime: 100 * time.Second,
	}
}

func MysqlInit(mysqlInfo *MysqlInfo) *sql.DB {
	dbDSN := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&%s", mysqlInfo.Username, mysqlInfo.Password,
		mysqlInfo.Network, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.Database, mysqlInfo.Charset, mysqlInfo.Other)
	fmt.Println("mysql info:", dbDSN)
	var mysqlDbErr error
	MysqlTw, mysqlDbErr = sql.Open("mysql", dbDSN)
	if mysqlDbErr != nil {
		// 打开连接失败
		panic("Incorrect data source configuration: " + mysqlDbErr.Error())
	}

	// 最大连接数
	MysqlTw.SetMaxOpenConns(mysqlInfo.MaxOpenConns)
	// 闲置连接数
	MysqlTw.SetMaxIdleConns(mysqlInfo.MaxIdleConns)
	// 最大连接周期，超过时间的连接就close
	MysqlTw.SetConnMaxLifetime(mysqlInfo.ConnMaxLifetime)

	if mysqlDbErr = MysqlTw.Ping(); nil != mysqlDbErr {
		panic("Database connection failed: " + mysqlDbErr.Error())
	}
	return MysqlTw
}

func MysqlClose() {
	MysqlTw.Close()
}
