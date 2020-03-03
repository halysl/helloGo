package main

import (
	"fmt"
	"log"
	"os"
)

func logDemo() {
	log.Println("这是第一条日志")
	fmt.Println("这是第一个数据")
	
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很详细的日志")

	log.SetPrefix("[测试用]")
	fmt.Println("现在的日志前缀是：", log.Prefix())
	log.Println("这是加了前缀的日志")
}

func logToFile() {
	logFile, err := os.OpenFile("/tmp/test.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := logFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.SetOutput(logFile)
	log.Println("存在文件里的日志")
}

func customLog() {
	logFile, err := os.OpenFile("/tmp/test.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := logFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	logger := log.New(logFile, "【自定义】", log.Llongfile|log.Lmicroseconds|log.Ldate)
	logger.Println("自定义输出的日志")
}