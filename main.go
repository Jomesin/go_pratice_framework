package main

import (
	conf "Go_workspace/conf"
	users "Go_workspace/src/users"
	logger "Go_workspace/utils/log"
	"net/http"
)

func registerRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	users.RegisterRoutes(mux) // 用户管理注册路由
	return mux
}

func startServer() error {
	address := conf.Cfg.Base.Address
	mux := registerRoutes() // 总注册路由执行函数
	return http.ListenAndServe(address, mux)
}

func main() {
	err := startServer()
	if err != nil {
		logger.Error("Error starting server")
	}
}
