package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"http_server/api/dbops"
	"http_server/api/defs"
	"http_server/api/session"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := defs.UserCredential{}
	// 如果参数解析错误
	if err := json.Unmarshal(res, ubody); err != nil {
		// 发送错误
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	// 如果新增用户到数据库错误
	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		// 发送错误
		sendErrorResponse(w, defs.ErrorDataBaseHandle)
		return
	}
	// 获取session_id
	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{
		Success:   false,
		SessionId: id,
	}
	if resp, err := json.Marshal(su); err != nil {
		// 发送错误
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}

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
