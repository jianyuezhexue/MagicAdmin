package magic

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// NewNameReplacer 库名，表名，字段名命名策略
type NewNameReplacer struct{}

// Replace 库名，表名，字段名替换
func (r NewNameReplacer) Replace(name string) string {
	if name == "" {
		return name
	}
	first := name[0:1]
	lowerFirst := strings.ToLower(first)
	return strings.Replace(name, first, lowerFirst, 1)
}

// initGorm 初始化gorm
func initGorm() *gorm.DB {
	config := Config.Mysql
	dsn := config.UserName + ":" + config.Password + "@tcp(" + config.Path + ")/" + config.DbName + "?" + config.Config
	mysqlConfig := mysql.Config{DSN: dsn}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             2 * time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Error,    // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,            // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
			NoLowerCase:   true, // 不使用小写
			NameReplacer:  NewNameReplacer{},
		},
	})
	if err != nil {
		panic(fmt.Errorf("链接数据库失败: %s", err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	return db
}
