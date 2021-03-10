package main

import (
	"log"
	"os"
	"time"
)

var (
	infoLog  *log.Logger
	errorLog *log.Logger
)

func printLog() {
	infoLog.Println("info")
	errorLog.Println("error")
}

func main() {
	//日志输出文件
	// file, err := os.Create("/var/log/1")
	// if err != nil {
	// 	log.Fatalln("Faild to open error logger file:", err)
	// }
	//自定义日志格式
	infoLog = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(os.Stdout, "Error:", log.Ldate|log.Ltime|log.Lshortfile)
	c := time.Tick(1 * time.Second)
	for range c {
		printLog()
	}
}
