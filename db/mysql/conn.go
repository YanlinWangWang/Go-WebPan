package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//默认初始化操作
func init(){
	//打开对应的数据库连接
	db, _ =sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
	//设置最大连接数
	db.SetMaxOpenConns(1000)

	err := db.Ping()

	if err!=nil{
		fmt.Println("连接失败 错误信息如下： "+ err.Error())
		os.Exit(1)
	}
}

// DBConn : 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}

