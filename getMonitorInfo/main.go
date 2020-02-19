package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/halysl/helloGO/getMonitorInfo/getInfo"
	"github.com/halysl/helloGo/getMonitorInfo/db"
)

func main() {

	db.InitDB()

	data := getInfo.GetOriginInfo()
	// insert_sql(data)
	for {
		time.Sleep(10 * time.Second)
		data = getInfo.GetOriginInfo()
		db.Update_sql(data)
	}
}
