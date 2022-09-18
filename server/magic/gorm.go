package magic

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// smallHump 库名，表名，字段名小驼峰策略
type smallHump struct{}

// Replace 库名，表名，字段名小驼峰替换策略
func (s smallHump) Replace(name string) string {
	if name == "" {
		return name
	}

	// // 兼容go-lint提醒
	// if name == "ID" {
	// 	return "id"
	// }
	// if name == "UUID" {
	// 	return "uuid"
	// }
	// if name == "URL" {
	// 	return "url"
	// }

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
	// 自定义打印
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold:             2 * time.Second, // 慢 SQL 阈值
	// 		LogLevel:                  logger.Error,    // 日志级别
	// 		IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
	// 		Colorful:                  true,            // 禁用彩色打印
	// 	},
	// )
	// orm配置
	ormConfig := &gorm.Config{
		// Logger: newLogger,
		Logger: logger.Default.LogMode(logger.Info), // 配置日志级别，打印出所有的sql
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,        // 使用单数表名
			NoLowerCase:   true,        // 不使用小写
			NameReplacer:  smallHump{}, // 小驼峰策略
		},
	}
	// 链接数据库
	db, err := gorm.Open(mysql.New(mysqlConfig), ormConfig)
	if err != nil {
		panic(fmt.Errorf("链接数据库失败: %s", err))
	}
	// DB设置
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
