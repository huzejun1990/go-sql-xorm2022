// @Author huzejun 2022/11/27 3:05:00
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

func main() {
	// 数据库连接基本信息
	var (
		userName  string = "root"
		password  string = "123456"
		idAddress string = "127.0.0.1"
		port      int    = 3306
		dbName    string = "go_test"
		charset   string = "utf8mb4"
	)

	// 构建数据库连接信息
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, idAddress, port, dbName, charset)
	//dataSourceName := fmt.Printf("%s:%s@tcp(%s:%d)/%s?charset=%s",userName,password,idAddress,port,dbName,charset)

	// xorm.NewEngine
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println("数据库连接失败！")
	}

	type User struct {
		Id      int64
		Name    string
		Age     int
		Passwd  string    `xorm:"varchar(200)"`
		Created time.Time `xorm:"created"`
		Updated time.Time `xorm:"updated"`
	}

	// engine.Sync
	err = engine.Sync(new(User))
	if err != nil {
		fmt.Println("表结构同步失败！")
	}
}
