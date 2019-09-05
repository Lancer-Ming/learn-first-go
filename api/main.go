package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandles() *httprouter.Router {
	router := httprouter.New()
	// 注册
	router.GET("/", CreateUser)

	// 登录
	router.POST("/user/:user_name", Login)
	return router
}


func main() {
	r := RegisterHandles()
	http.ListenAndServe(":8001", r)
}
