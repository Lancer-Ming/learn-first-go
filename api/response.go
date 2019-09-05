package main

import (
	"encoding/json"
	"http_server/api/defs"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	w.WriteHeader(errResp.HttpSC)

	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))

}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}