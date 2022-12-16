package magic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/service"
)

// 接口返回成功
func HttpSuccess(c *gin.Context, res service.BackData) {
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
	// 获取开始时间
	start := c.GetInt64("magicStartTime")

	// 计算耗时
	end := time.Now().UnixNano()
	costTime, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", (float64(end)-float64(start))/1000000), 64)

	// 返回结果
	res := service.BackData{Code: code, Msg: msg, Data: data, CostTime: costTime}

	// 预埋操作记录参数
	c.Set("magicMsg", msg)
	recodeBytes, _ := json.Marshal(res)
	c.Set("magicResp", string(recodeBytes))

	c.JSON(http.StatusOK, res)
}
