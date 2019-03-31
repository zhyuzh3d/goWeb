package util

import (
	"app/uds"
	"encoding/json"
	"net/http"
)

//WWrite 向用户返回信息,返回resp和一个错误信息
func WWrite(w http.ResponseWriter, code int, msg string, data interface{}) (uds.Respons, error) {
	resp := uds.Respons{Code: code, Msg: msg, Data: data}
	var err error
	dt, err1 := json.Marshal(resp)
	if err1 != nil {
		err = err1
	}
	_, err2 := w.Write([]byte(string(dt)))
	if err2 != nil {
		err = err2
	}
	return resp, err
}
