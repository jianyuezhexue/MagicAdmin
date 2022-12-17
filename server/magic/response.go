package magic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 返回结构
type BackData struct {
	Code     int     `json:"code"`     // 返回吗
	Msg      string  `json:"msg"`      // 返回消息
	Data     any     `json:"data"`     // 返回数据
	CostTime float64 `json:"costTime"` // 接口耗时
}

// 返回上一层数据|带行号
func Back(code int, msg string, data any) BackData {
	_, _, line, _ := runtime.Caller(1)
	if code != 0 {
		msg = msg + "[" + strconv.Itoa(line) + "]"
	}
	return BackData{code, msg, data, 0}
}

// 接口返回成功
func HttpSuccess(c *gin.Context, res BackData) {
	// 获取开始时间
	start := c.GetInt64("magicStartTime")

	// 计算耗时
	end := time.Now().UnixNano()
	costTime, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", (float64(end)-float64(start))/1000000), 64)

	// 补全耗时
	res.CostTime = costTime

	// 预埋操作记录参数
	c.Set("magicMsg", res.Msg)
	recodeBytes, _ := json.Marshal(res.Data)
	c.Set("magicResp", string(recodeBytes))

	c.JSON(http.StatusOK, res)
}

// 接口返回失败
func HttpFail(c *gin.Context, code int, msg string, data any) {
	_, _, line, _ := runtime.Caller(1)
	msg = msg + "[" + strconv.Itoa(line) + "]"

	// 获取开始时间
	start := c.GetInt64("magicStartTime")

	// 计算耗时
	end := time.Now().UnixNano()
	costTime, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", (float64(end)-float64(start))/1000000), 64)

	// 返回结果
	res := BackData{Code: code, Msg: msg, Data: data, CostTime: costTime}

	// 预埋操作记录参数
	c.Set("magicMsg", msg)
	recodeBytes, _ := json.Marshal(res)
	c.Set("magicResp", string(recodeBytes))

	c.JSON(http.StatusOK, res)
}
