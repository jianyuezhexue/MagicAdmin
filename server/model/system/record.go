package system

import (
	"time"

	"github.com/jianyuezhexue/MagicAdmin/model"
)

// 如果含有time.Time 请自行import time包
type SysOperationRecord struct {
	model.BaseOrm
	Ip           string        `json:"ip"`            // 请求ip
	Method       string        `json:"method"`        // 请求方法
	Path         string        `json:"path"`          // 请求路径
	Status       int           `json:"status"`        // 请求状态
	Latency      time.Duration `json:"latency"`       // 延迟
	Agent        string        `json:"agent"`         // 代理
	ErrorMessage string        `json:"error_message"` // 错误信息
	Body         string        `json:"body"`          // 请求Body
	Resp         string        `json:"resp"`          // 响应Body
	UserId       int           `json:"userId"`        // 用户id
	User         User          `json:"user"`
}
