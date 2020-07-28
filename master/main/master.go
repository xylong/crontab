package main

import (
	"crontab/master"
	"flag"
	"fmt"
	"runtime"
)

var (
	// 配置文件路径
	confFile string
)

// initArgs 解析命令行参数
func initArgs() {
	flag.StringVar(&confFile, "config", "master.json", "指定master.json")
	flag.Parse()
}

// initEnv 初始化线程数
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var err error

	// 初始化命令行参数
	initArgs()
	// 初始化线程
	initEnv()
	// 加载配置
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}
	// 任务管理器
	if err = master.InitJobManager(); err != nil {
		goto ERR
	}
	// 启动api服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	return
ERR:
	fmt.Println(err)
}
