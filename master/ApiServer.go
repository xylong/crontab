package master

import (
	"net"
	"net/http"
	"time"
)

var (
	Api *ApiServer
)

// ApiServer 任务http接口
type ApiServer struct {
	httpServer *http.Server
}

// handleJobSave 保存任务接口
func handleJobSave(w http.ResponseWriter,r *http.Request)  {
	
}

func InitApiServer() error {
	// 路由
	mux:=http.NewServeMux()
	mux.HandleFunc("job/save",handleJobSave)

	// tcp监听
	listener,err :=net.Listen("tcp",":8070")
	if err!=nil {
		return err
	}

	// http服务
	httpServer:=http.Server{
		ReadTimeout:       time.Second*2,
		WriteTimeout:      time.Second*2,
		Handler: mux,
	}

	Api=&ApiServer{httpServer: &httpServer}
	go httpServer.Serve(listener)

	return nil
}