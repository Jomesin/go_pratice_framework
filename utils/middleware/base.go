package middleware

import (
	"Go_workspace/conf"
	logger "Go_workspace/utils/log"
	"net/http"
)

type requestBefore func(r *http.Request, next http.Handler) http.Handler // 定义请求之前中间件
type responseAfter func(w http.ResponseWriter)                           // 定义响应后中间件

var (
	PublicBeforeMiddlewares []requestBefore = []requestBefore{
		LoginMiddlewareBefore,
		PersonMiddlewareBefore,
	}
	PublicAfterMiddlewares []responseAfter = []responseAfter{
		OperationMiddlewareAfter,
	}
	sem = make(chan struct{}, conf.Cfg.Middleware.MaxConcurrentRequests)
)

func LoadMiddleware(middleBefore []requestBefore, middleAfter []responseAfter, next http.Handler) http.Handler {
	// 公用中间件集中加载
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 限制同时访问数量
		select {
		case sem <- struct{}{}:
			defer func() { <-sem }() // 释放信号量
			defer func() {
				if r := recover(); r != nil {
					http.Error(w, "内部服务器错误", http.StatusInternalServerError)
				}
			}()
			// 加载请求前的中间件
			middleBe := len(middleBefore)
			for i := 0; i < middleBe; i++ {
				next = middleBefore[i](r, next)
			}
			next.ServeHTTP(w, r)
			// 加载响应后的中间件
			middleAf := len(middleAfter)
			for i := 0; i < middleAf; i++ {
				middleAfter[i](w)
			}
		default:
			logger.Info("限制请求数量")
			http.Error(w, "请求过多,请稍后再试,或访问其他服务", http.StatusTooManyRequests)
		}

	})
}
