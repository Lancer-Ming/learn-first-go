package main

import (
	"github.com/julienschmidt/httprouter"
	"http_server/api/defs"
	"net/http"
)

// 创建一个中间件的struct
type middlewareHandle struct {
	r *httprouter.Router
}

// 用函数来接收存储httprouter.Router
func newMiddlewareHandle(r *httprouter.Router) middlewareHandle {
	m := middlewareHandle{}
	m.r = r
	return m
}

// 重写HttpServer
func (m middlewareHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	if !validateUserSession(r) {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
	}
	// 然后再执行httprouter.Router的ServeHTTP
	m.r.ServeHTTP(w, r)
}

func RegisterHandles() *httprouter.Router {
	router := httprouter.New()
	// 注册
	router.POST("/", CreateUser)

	// 登录
	router.POST("/user/:user_name", Login)
	return router
}


func main() {
	r := RegisterHandles()
	m := newMiddlewareHandle(r)
	http.ListenAndServe(":8001", m)
}

