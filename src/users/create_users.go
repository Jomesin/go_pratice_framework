package users

import (
	logger "Go_workspace/utils/log"
	utils "Go_workspace/utils/responses"
	"net/http"
)

type Response struct {
	Data    []interface{}
	Message string
	Status  int
}

func CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.URL.Path)
	dataArray := make([]interface{}, 5, 10)
	utils.MakeResponse(w, utils.OK, dataArray)
}
