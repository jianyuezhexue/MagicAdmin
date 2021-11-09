package main

import (
	"fmt"

	"github.com/jianyuezhexue/MagicAdmin/magic"
)

func main() {
	magic.Logger.Info("测试测试测试")
	// sql := `CREATE TABLE datareport (
	// 	openId varchar(64) NOT NULL COMMENT '用户id',
	// 	accType char(2) NOT NULL COMMENT '登录类型',
	// 	createTime datetime DEFAULT NULL COMMENT '创建小屋时间',
	// 	shareTime datetime DEFAULT NULL COMMENT '分享时间',
	// 	PRIMARY KEY (openId,accType) USING BTREE
	//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`

	// magic.Orm.Exec(sql)
	// magic.Orm.Exec(sql)
	fmt.Print(magic.Redis)
}
