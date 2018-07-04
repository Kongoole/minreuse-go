package service

type Response struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

const (
	HTTP_OK           = 200
	HTTP_SERVER_ERROR = 500
)
