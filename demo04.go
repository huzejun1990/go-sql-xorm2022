// @Author huzejun 2022/11/27 19:14:00
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

	// 查询
	/*	results,err := engine.Query("select * from user")
		fmt.Println(results)
		results2,err := engine.QueryString("select * from user")
		fmt.Println(results2)
		results3,err := engine.QueryInterface("select * from user")
		fmt.Println(results3)*/

	// Get	// SELECT * FROM user limit 1
	user := User{}
	engine.Get(&user)
	fmt.Println(user)

	// 指定条件来查询用户
	user1 := User{Name: "dream"}
	engine.Where("name=?", user1.Name).Desc("id").Get(&user1)
	fmt.Println(user1)

	// 获取指定字段的值
	var name string
	engine.Table(&user).Where("id =10001").Cols("passwd").Get(&name)
	fmt.Println(name)

	// Find 查询多条记录
	var users []User
	engine.Where("passwd=123456").And("age=18").Limit(10, 0).Find(&users)
	fmt.Println(users)

	// Count 记录总条数
	user = User{Age: 18}
	counts, err := engine.Count(&user)
	fmt.Println(counts)

	// Iterate 和 Rows 根据条件遍历数据库
	/*	engine.Iterate(&User{Passwd: "123456"}, func(idx int, bean interface{}) error {
		user := bean.(*User)
		fmt.Println(user)
		return nil
	})*/

	rows, err := engine.Rows(&User{Passwd: "123456"})
	defer rows.Close()
	userBean := new(User)
	for rows.Next() {
		rows.Scan(userBean)
		fmt.Println(userBean)
	}

}
