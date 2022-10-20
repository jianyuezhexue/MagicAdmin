package service

import (
	"runtime"
	"strconv"
)

// 返回结构
type BackData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// 返回上一层数据|带行号
func Back(code int, msg string, data any) BackData {
	_, _, line, _ := runtime.Caller(1)
	if code != 0 {
		msg = msg + "[" + strconv.Itoa(line) + "]"
	}
	return BackData{code, msg, data}
}
