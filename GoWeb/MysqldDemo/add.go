package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:hou@tcp(localhost:3306)/GO")
	_ = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
	}

	defer func() {
		if db != nil {
			_ = db.Close()
		}
		fmt.Println("关闭连接")
	}()

	stmt, err := db.Prepare("insert into people values(default,?,?)")
	if err != nil {
		fmt.Println("预处理失败", err)
	}
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()

	res, err := stmt.Exec("xiaopang", "shenzhen")

	if err != nil {
		fmt.Println("执行SQL出现错误")
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("获取主键失败", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取结果失败", err)
	}

	fmt.Println(id, count)

	if count > 0 {
		fmt.Println("新增成功")
	} else {
		fmt.Println("新增失败")
	}
}
