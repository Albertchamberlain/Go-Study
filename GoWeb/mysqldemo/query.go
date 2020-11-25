package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:hou@tcp(localhost:3306)/GO")
	db.Ping()

	if err != nil {
		fmt.Println("数据库连接失败")
	}

	defer func() {
		if db != nil {
			db.Close()
			fmt.Println("关闭连接")
		}
	}()

	stmt, err := db.Prepare("select * from people")
	if err != nil {
		fmt.Println("预处理失败", err)
	}

	defer func() {
		if stmt != nil {
			stmt.Close()
			fmt.Println("stmt关闭")
		}
	}()
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("查询失败", err)
	}
	for rows.Next() {
		var id int
		var name, address string
		rows.Scan(&id, &name, &address)
		fmt.Println(id, name, address)
	}
	defer func() {
		if rows != nil {
			rows.Close()
			fmt.Println("关闭结果集")
		}
	}()
}
