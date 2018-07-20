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
	HTTP_BADPARAMS = 100001
)

// JSONResp sends json response to client
func (r Response) JSONResponse(w http.ResponseWriter) {
	resp, _ := json.Marshal(r)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(resp)
}
