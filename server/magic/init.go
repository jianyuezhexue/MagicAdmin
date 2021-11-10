package magic

import (
	"fmt"

	"github.com/jianyuezhexue/MagicAdmin/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// Config 配置
	Config config.Config
	// Logger 日志
	Logger *zap.Logger
	// Orm ORM
	Orm *gorm.DB
	// Redis redisClient
	Redis RedisClient
)

// 读取配置
func initConfig() {
	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
	}
}

// 初始化配置
func init() {
	// 读取配置
	initConfig()
	// 初始化日志
	Logger = initZap()
	// 初始化ORM
	Orm = initGorm()
	// 初始化Redis
	Redis = NewRedisCache()
}

// 启动服务
func run() {
	// 中间件，路由，启动服务
}
