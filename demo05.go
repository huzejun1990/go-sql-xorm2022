// @Author huzejun 2022/11/27 20:52:00
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

	session := engine.NewSession()
	defer session.Close()

	session.Begin()
	defer func() {
		err := recover()
		if err != nil {
			// 回滚
			fmt.Println(err)
			fmt.Println("Rollback")
			session.Rollback()
		} else {
			session.Commit()
		}
	}()

	user1 := User{Id: 10007, Name: "dream1111", Age: 18, Passwd: "1234213"}
	if _, err := session.Insert(&user1); err != nil {
		panic(err)
	}

	user2 := User{Name: "dream2222", Age: 3}
	if _, err := session.Where("id=1000").Update(&user2); err != nil {
		panic(err)
	}

	if _, err := session.Exec("delete from user where name = 'dream1111'"); err != nil {
		panic(err)
	}
}
