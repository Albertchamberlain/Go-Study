package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:hou@tcp(localhost:3306)/GO")
	db.Ping()
	//Error处理
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	//关闭连接
	defer func() {
		if db != nil {
			db.Close()
			fmt.Println("关闭连接")
		}
	}()
	stmt, err := db.Prepare("delete from people where id=?")
	//错误处理
	if err != nil {
		fmt.Println("预处理失败", err)
	}
	//关闭对象
	defer func() {
		if stmt != nil {
			stmt.Close()
			fmt.Println("stmt关闭")
		}
	}()

	res, err := stmt.Exec(1)
	//错误处理
	if err != nil {
		fmt.Println("执行SQL出现错误")
	}
	//受影响行数
	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取结果失败", err)
	}
	if count > 0 {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}
}
