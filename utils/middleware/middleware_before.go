package middleware

import (
	logger "Go_workspace/utils/log"
	"net/http"
)

func LoginMiddlewareBefore(r *http.Request, next http.Handler) http.Handler {
	// 登录中间件
	logger.Info("测试登录前")
	return next
}

func PersonMiddlewareBefore(r *http.Request, next http.Handler) http.Handler {
	logger.Info("用户处理数据")
	return next
}
