package main

import (
	"mycode/getMonitorInfo/db"
	"mycode/getMonitorInfo/getInfo"
	"time"
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

