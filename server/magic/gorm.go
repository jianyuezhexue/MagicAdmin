package magic

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormLogger struct
type GormLogger struct{}

// Printf - Log Formatter
func (*GormLogger) Printf(msg string, data ...interface{}) {
	// fmt.Print(data[0]) // 报错文件路径
	// fmt.Print(data[1]) // 报错信息
	// fmt.Print(data[2]) // 消耗时间
	// fmt.Print(data[3]) // 影响行数
	// fmt.Print(data[4]) // SQL语句
	// fmt.Printf(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

// initGorm 初始化gorm
func initGorm() *gorm.DB {
	config := Config.Mysql
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Path + ")/" + config.Dbname + "?" + config.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		// &GormLogger{},
		logger.Config{
			SlowThreshold:             2 * time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Error,    // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,            // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(fmt.Errorf("链接数据库失败: %s", err))
	}
	// TODO:设置日志打印
	// db.Logger.LogMode()

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	return db
}