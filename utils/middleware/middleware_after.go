package middleware

import (
	logger "Go_workspace/utils/log"
	"net/http"
)

func OperationMiddlewareAfter(w http.ResponseWriter) {
	logger.Info("响应之后处理")
}
