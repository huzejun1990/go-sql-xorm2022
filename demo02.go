// @Author huzejun 2022/11/27 5:20:00
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

	type User struct {
		Id      int64
		Name    string
		Age     int
		Passwd  string    `xorm:"varchar(200)"`
		Created time.Time `xorm:"created"`
		Updated time.Time `xorm:"updated"`
	}

	// engine.Insert 插入 一个对象，返回值：受影响的行数
	user := User{Id: 10000, Name: "lucy", Age: 18, Passwd: "123456"}
	n, _ := engine.Insert(&user)
	fmt.Println(n)
	if n >= 1 {
		fmt.Println("数据插入成功！")
	}

	user1 := User{Id: 10001, Name: "lucy1", Age: 18, Passwd: "123456"}
	user2 := User{Id: 10002, Name: "lucy3", Age: 18, Passwd: "123456"}
	n, _ = engine.Insert(&user1, &user2)
	fmt.Println(n)
	if n >= 1 {
		fmt.Println("数据插入成功！")
	}

	var users []User
	users = append(users, User{Id: 10003, Name: "lucy2", Age: 18, Passwd: "123456"})
	users = append(users, User{Id: 10004, Name: "lucy4", Age: 18, Passwd: "123456"})
	n, _ = engine.Insert(users)
}
