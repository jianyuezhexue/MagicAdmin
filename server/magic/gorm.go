package magic

import (
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// smallHump 库名，表名，字段名小驼峰策略
type smallHump struct{}

// todo这里要修改

// Replace 库名，表名，字段名小驼峰替换策略
func (s smallHump) Replace(name string) string {
	if name == "" {
		return name
	}

	// 兼容go-lint提醒
	if name == "ID" {
		return "id"
	}
	if name == "UUID" {
		return "uuid"
	}
	if name == "URL" {
		return "url"
	}

	// 首字母小写
	newName := ""
	first := name[0:1]
	lowerFirst := strings.ToLower(first)
	newName = lowerFirst + name[1:len(name)-0]

	// 尾字母D转小写
	last := name[len(name)-1 : len(name)-0]
	if last == "D" {
		newName = newName[0:len(newName)-1] + "d"
	}

	return newName
}

// initGorm 初始化gorm
func initGorm() *gorm.DB {
	// 组合dsn
	config := Config.Mysql
	dsn := config.Dsn()
	mysqlConfig := mysql.Config{DSN: dsn}

	// orm配置
	ormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,        // 使用单数表名
			NoLowerCase:   true,        // 不使用小写
			NameReplacer:  smallHump{}, // 小驼峰策略
		},
	}
	// 链接数据库
	db, err := gorm.Open(mysql.New(mysqlConfig), ormConfig)
	if err != nil {
		Print("链接数据库失败")
		return nil
	}
	Print("链接数据库成功")

	// DB设置
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)

	// 查询和设置自动断开链接的时间
	// SHOW VARIABLES LIKE '%timeout%';
	// wait_timeout = 120
	sqlDB.SetConnMaxLifetime(time.Second * 100)
	return db
}
