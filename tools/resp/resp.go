package resp

import (
	"encoding/json"
	"net/http"
)

// JSONResult 返回结果
type JSONResult struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data map[string]string `json:"data"`
}

// Success 返回成功消息
func Success(w http.ResponseWriter) {
	Response(w, 200, "请求成功")
}

// ErrorParam 返回参数错误
func ErrorParam(w http.ResponseWriter, err error) {
	Response(w, 400, "参数错误")
}

// ErrorAuth 权限错误
func ErrorAuth(w http.ResponseWriter) {
	Response(w, 401, "请登录")
}

// ErrorServer 返回服务器错误
func ErrorServer(w http.ResponseWriter, err error) {
	Response(w, 500, "服务器错误")
}

// Response 返回结果
func Response(w http.ResponseWriter, code int, content string) {
	var msg []byte
	msg, _ = json.Marshal(JSONResult{Code: code, Msg: content})
	w.WriteHeader(code)
	w.Write(msg)
}

// RespData 返回结果
func RespData(w http.ResponseWriter, code int, content string, data map[string]string) {
	var msg []byte
	msg, _ = json.Marshal(JSONResult{Code: code, Msg: content, Data: data})
	w.WriteHeader(code)
	w.Write(msg)
}
