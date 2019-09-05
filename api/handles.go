package main

import (
	"github.com/julienschmidt/httprouter"
	"http_server/api/dbops"
	"io"
	"net/http"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Hello, I begin to learn golang language!!!!!")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("user_name")
	io.WriteString(w, username)
}

func CreateComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid, _ := strconv.Atoi(p.ByName("uid"))
	vid := p.ByName("vid")
	content := p.ByName("content")
	err := dbops.AddComment(uid, vid, content)
	if err != nil {
	}
}
