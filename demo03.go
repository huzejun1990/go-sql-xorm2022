// @Author huzejun 2022/11/27 19:00:00
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

	// xorm.NewEngine
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println("数据库连接失败！")
	}

	// 构建插入的用户数据
	type User struct {
		Id      int64
		Name    string
		Age     int
		Passwd  string    `xorm:"varchar(200)"`
		Created time.Time `xorm:"created"`
		Updated time.Time `xorm:"updated"`
	}

	// .Update(&user)
	user := User{Name: "lucy", Age: 28}
	n, _ := engine.ID(1000).Update(&user)
	fmt.Println(n)

	// .Delete(&user)
	user = User{Name: "lucy"}
	n, _ = engine.ID(10000).Delete(&user)
	fmt.Println(n)

	engine.Exec("update user set age = ? where id = ? ", 10, 10003)

}
