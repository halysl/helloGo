package db

import (
	"database/sql"
	"log"
	"strings"
	"syscall"

	"github.com/halysl/hellogo/getMonitorInfo/static"

)

// DB 数据库连接池
var DB *sql.DB

// InitDB 注意方法名大写，就是 public
func InitDB() {

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{static.UserName, ":", static.Password, "@tcp(", static.Ip, ":", static.Port, ")/", static.DbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		log.Printf("opon database fail!!!")
		log.Fatal(err)
		syscall.Exit(-1)
	}
	log.Printf("connnect success!\n")
}

func execSQL(sql string) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal(err)
		log.Printf("[ERR]tx fail\n")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Fatal(err)
		log.Printf("[ERR]Prepare fail\n")
		return false
	}
	//将参数传递到sql语句中并且执行
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
		log.Printf("[ERR]Exec fail\n")
		return false
	}
	//将事务提交
	tx.Commit()
	log.Printf("Execute Sql Success!")
	return true
}
