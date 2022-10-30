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

// 初始化定时任务
func InitCron() {
	c := cron.New()
	// 每三秒保存一次记录
	c.AddFunc("*/3 * * * * *", storeRecode)
	c.Start()
}
