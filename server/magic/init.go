package magic

import (
	"fmt"

	"github.com/jianyuezhexue/MagicAdmin/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// Config 全局配置变量
	Config config.Config
	// Logger 日志
	Logger *zap.Logger
	// Orm gorm
	Orm *gorm.DB
)

// 初始化配置
func init() {
	// 读取配置
	initConfig()

	// 初始化日志
	Logger = initZap()

	// 初始化ORM

	// 初始化Redis

	// 初始化DB
}

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

// 启动服务
func run() {

}
