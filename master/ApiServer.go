package master

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

// ApiServer api服务
var ApiServer *apiServer

// ApiServer 任务http接口
type apiServer struct {
	httpServer *http.Server
}

// handleJobSave 保存任务接口
func handleJobSave(w http.ResponseWriter, r *http.Request) {

}

// InitApiServer 初始化服务
func InitApiServer() (err error) {
	var (
		mux      *http.ServeMux
		listener net.Listener
	)
	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)
	// 启动监听
	if listener, err = net.Listen("tcp", ":"+strconv.Itoa(Conf.ApiPort)); err != nil {
		return
	}
	// 创建http服务
	server := &http.Server{
		ReadTimeout:  time.Duration(Conf.ApiReadTimeOut) * time.Millisecond,
		WriteTimeout: time.Duration(Conf.ApiWriteTimeOut) * time.Millisecond,
		Handler:      mux,
	}
	// 单例
	ApiServer = &apiServer{
		httpServer: server,
	}
	// 启动服务
	go server.Serve(listener)

	return
}
