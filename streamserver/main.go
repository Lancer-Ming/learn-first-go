package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video/:video-id", StreamVideoHandler)
	router.POST("/video/:video-id", UploadVideoHandler)

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":9000", r)
}