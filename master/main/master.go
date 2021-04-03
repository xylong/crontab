package main

import (
	"crontab/master"
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	if err:=master.InitApiServer();err!=nil {
		fmt.Println(err.Error())
	}
}
