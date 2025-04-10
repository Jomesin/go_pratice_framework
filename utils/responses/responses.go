package responses

import (
	logger "Go_workspace/utils/log"
	"encoding/json"
	"net/http"
	"time"
)

type response struct {
	Data    []interface{}
	Message string
	Status  int
}

const (
	OK = 0
)

var resMsg = map[int]string{
	OK: "success",
}

func MakeResponse(w http.ResponseWriter, resCode int, dataArray []interface{}) {
	val, exists := resMsg[resCode]
	if !exists {
		val = resMsg[OK]
	}

	// 获取当前时间
	currentTime := time.Now().Local()
	// 将当前时间格式化为 HTTP 日期格式
	dateHeader := currentTime.Format(time.RFC1123)

	res := response{Message: val, Status: resCode, Data: dataArray}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST")
	w.Header().Set("Accept-Language", "zh-CN,zh;q=0.9")
	w.Header().Set("Date", dateHeader)

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		logger.Info("Serialization failed")
	}
}
