package middle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
)

// 记录操作记录
func Record() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接口开始时间
		start := time.Now().UnixNano()
		c.Set("magicStartTime", start)

		claims := magic.TokenInfo(c)

		// 收集参数
		var record system.RecordBase
		record.Ip = c.ClientIP()             // IP
		record.UserId = int(claims.ID)       // 用户ID
		record.UserName = claims.UserName    // 用户昵称
		record.Agent = c.Request.UserAgent() // Agent
		record.Method = c.Request.Method     // Method
		record.Path = c.Request.URL.Path     // Path
		record.Params = getAllParams(c)      // Params

		c.Next()

		// 计算耗时
		end := time.Now().UnixNano()
		costTime, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", (float64(end)-float64(start))/1000000), 64)

		record.Status = c.Writer.Status()      // 返回状态
		record.Msg = c.GetString("magicMsg")   // 返回信息
		record.CostTime = costTime             // 接口耗时
		record.Resp = c.GetString("magicResp") // 接口返回

		// 删除记录
		c.Set("magicMsg", "")
		c.Set("magicResp", "")

		// 查询记录不存储返回
		if record.Method == http.MethodGet && record.Path == "/recode" {
			record.Resp = "不记录" // 接口返回
		}

		// 设置熔断等级
		// todo

		// 发送到队列
		recordBytes, _ := json.Marshal(record)
		magic.Redis.Lpush(magic.RecodeListKey, string(recordBytes))
	}
}

// 获取所有类型的参数
func getAllParams(c *gin.Context) string {
	var body string

	// GET 请求
	if c.Request.Method == http.MethodGet {
		query := c.Request.URL.Query()
		var queryMap = make(map[string]any, len(query))
		for k := range query {
			queryMap[k] = c.Query(k)
		}
		bytes, _ := json.Marshal(queryMap)
		body = string(bytes)
		return body
	}

	// DELETE 请求
	if c.Request.Method == http.MethodDelete {
		m := make(map[string]string)
		for _, v := range c.Params {
			m[v.Key] = v.Value
		}

		bytes, _ := json.Marshal(m)
		body = string(bytes)
		return body
	}

	// POST,PUT,PATCH 请求
	bodyBytes, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	body = string(bodyBytes)

	return body
}
