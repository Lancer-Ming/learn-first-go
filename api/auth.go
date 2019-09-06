package main

import (
	"http_server/api/session"
	"net/http"
)


var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"
func validateUserSession(r *http.Request) bool {
	// 从headers 获取 sid
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	// 查看是否过期
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func validateUser(r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		return false
	}
	return true
}