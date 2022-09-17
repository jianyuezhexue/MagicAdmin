package magic

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
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
	// SingleFlight 防止缓存击穿
	SingleFlight = &singleflight.Group{}
	// LocalCache 本地缓存
	LocalCache local_cache.Cache
	// once
	once sync.Once
)

// 读取配置
func initConfig() {
	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置信息错误: %s", err))
	}
	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
	}
}

// 初始化配置
func init() {
	// 读取配置
	once.Do(func() { initConfig() })
	// 初始化日志
	Logger = initZap()
	// 初始化ORM
	Orm = initGorm()
	// 初始化Redis
	Redis = NewRedisCache()
	fmt.Println("-------------------------- 初始化变量完成 --------------------------")
}

// Server Server
type Server interface {
	ListenAndServe() error
}

// InitServer 配置server
func InitServer(address string, router *gin.Engine) Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
