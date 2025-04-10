package users

import (
	middle "Go_workspace/utils/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/users", middle.LoadMiddleware(middle.PublicBeforeMiddlewares, middle.PublicAfterMiddlewares, http.HandlerFunc(CreateUsersHandler)))
}
