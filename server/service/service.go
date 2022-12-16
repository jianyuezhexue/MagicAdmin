package service

// // 返回结构
// type BackData struct {
// 	Code     int     `json:"code"`     // 返回吗
// 	Msg      string  `json:"msg"`      // 返回消息
// 	Data     any     `json:"data"`     // 返回数据
// 	CostTime float64 `json:"costTime"` // 接口耗时
// }

// // 返回上一层数据|带行号
// func Back(code int, msg string, data any) BackData {
// 	_, _, line, _ := runtime.Caller(1)
// 	if code != 0 {
// 		msg = msg + "[" + strconv.Itoa(line) + "]"
// 	}
// 	return BackData{code, msg, data, 0}
// }
