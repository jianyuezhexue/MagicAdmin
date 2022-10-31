package schedule

import (
	"encoding/json"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
	"github.com/robfig/cron"
)

type Schedule struct{}

var ScheduleApp = new(Schedule)

// 每三秒保存一次记录
func storeRecode() {
	res, err := magic.Redis.Rpop(magic.RecodeListKey)
	if err == nil && res != "" {
		var record system.Record
		err = json.Unmarshal([]byte(res), &record)
		if err != nil {
			magic.Print(err.Error())
			return
		}

		// 记录入库
		err = serviceSystem.RecodeApp.Create(record)
		if err != nil {
			magic.Logger.Info("日志记录失败:" + err.Error())
		}
	}
}

// 每天的凌晨3点，删除15天前的日志记录
// todo

// 初始化定时任务
func InitCron() {
	c := cron.New()
	// 每三秒保存一次记录
	c.AddFunc("*/3 * * * * *", storeRecode)

	// 每天凌晨3点删除15天前的日志记录 | todo
	// c.AddFunc("0 0 3 * * ?", storeRecode)

	c.Start()
}

// 常用示例
// 每隔5秒执行一次：*/5 * * * * ?
// 每隔1分钟执行一次：0 */1 * * * ?
// 每天23点执行一次：0 0 23 * * ?
// 每天凌晨1点执行一次：0 0 1 * * ?
// 每月1号凌晨1点执行一次：0 0 1 1 * ?
// 在26分、29分、33分执行一次：0 26,29,33 * * * ?
// 每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?
