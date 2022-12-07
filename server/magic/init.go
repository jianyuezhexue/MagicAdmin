package magic

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/config"
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
	LocalCache *bigcache.BigCache
	// casbin
	cachedEnforcer *casbin.CachedEnforcer
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

// 初始化本地缓存
func initLocalCache() *bigcache.BigCache {
	localCache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		panic(fmt.Errorf("初始化本地缓存失败: %s", err))
	}
	return localCache
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
	// 初始化本地缓存
	LocalCache = initLocalCache()
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
