package service

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

const (
	HTTP_OK           = 200
	HTTP_SERVER_ERROR = 500
)

func (r Response) JSONResponse(w http.ResponseWriter) {
	result, _ := json.Marshal(r)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(result)
}
