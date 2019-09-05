package main

import (
	"encoding/json"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter) interface{} {
	data, err := json.Marshal(w)
	if err != nil {
		panic("Json stringify data failed")
	}
	return data
}

func sendNormalResponse(w http.ResponseWriter) {

}