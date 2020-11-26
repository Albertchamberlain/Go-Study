package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:hou@tcp(localhost:3306)/GO")
	_ = db.Ping()
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

	stmt, err := db.Prepare("update people set name=?,address=? where id=?")

	if err != nil {
		fmt.Println("预处理失败", err)
	}

	defer func() {
		if stmt != nil {
			_ = stmt.Close()
			fmt.Println("stmt关闭")
		}
	}()

	res, err := stmt.Exec("xiaohong", "shenzhen", 1)

	if err != nil {
		fmt.Println("执行SQL出现错误")
	}

	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取结果失败", err)
	}
	if count > 0 {
		fmt.Println("修改成功")
	} else {
		fmt.Println("修改失败")
	}
}
