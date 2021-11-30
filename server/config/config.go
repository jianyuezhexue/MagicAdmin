package config

// System 系统图配置
type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                               // 环境值
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                            // 端口值
	DbType        string `mapstructure:"dbType" json:"dbType" yaml:"dbType"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"ossType" json:"ossType" yaml:"ossType"`                   // Oss类型
	UseMultipoint bool   `mapstructure:"useMultipoint" json:"useMultipoint" yaml:"useMultipoint"` // 多点登录拦截
	LimitCountIP  int    `mapstructure:"iplimitCount" json:"iplimitCount" yaml:"iplimitCount"`
	LimitTimeIP   int    `mapstructure:"iplimitTime" json:"iplimitTime" yaml:"iplimitTime"`
}

// JWT jwt配置
type JWT struct {
	SigningKey  string `mapstructure:"signingKey" json:"signingKey" yaml:"signingKey"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expiresTime" json:"expiresTime" yaml:"expiresTime"` // 过期时间
	BufferTime  int64  `mapstructure:"bufferTime" json:"bufferTime" yaml:"bufferTime"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                // 签发者
}

// Zap 日志配置
type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                         // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                      // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                      // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`               // 日志文件夹
	ShowLine      bool   `mapstructure:"showLine" json:"showLine" yaml:"showLine"`                // 显示行
	EncodeLevel   string `mapstructure:"encodeLevel" json:"encodeLevel" yaml:"encodeLevel"`       // 编码级
	StacktraceKey string `mapstructure:"stacktraceKey" json:"stacktraceKey" yaml:"stacktraceKey"` // 栈名
	LogInConsole  bool   `mapstructure:"logInConsole" json:"logInConsole" yaml:"logInConsole"`    // 输出控制台
}

// Redis Redis配置
type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
}

// Mysql mysql配置
type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                         // 服务器地址:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                   // 高级配置
	DbName       string `mapstructure:"dbName" json:"dbname" yaml:"dbName"`                   // 数据库名
	UserName     string `mapstructure:"userName" json:"userName" yaml:"userName"`             // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`             // 数据库密码
	MaxIdleConns int    `mapstructure:"maxidleConns" json:"maxIdleConns" yaml:"maxidleConns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"maxOpenConns" json:"maxOpenConns" yaml:"maxOpenConns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"logMode" json:"logMode" yaml:"logMode"`                // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"logZap" json:"logZap" yaml:"logZap"`                   // 是否通过zap写入日志文件
}

// Dsn 组合链接参数
func (m *Mysql) Dsn() string {
	return m.UserName + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.DbName + "?" + m.Config
}

// Timer 定时器配置
type Timer struct {
	Start  bool     `mapstructure:"start" json:"start" yaml:"start"` // 是否启用
	Spec   string   `mapstructure:"spec" json:"spec" yaml:"spec"`    // CRON表达式
	Detail []Detail `mapstructure:"detail" json:"detail" yaml:"detail"`
}

// Detail 定时清除表配置
type Detail struct {
	TableName    string `mapstructure:"tableName" json:"tableName" yaml:"tableName"`          // 需要清理的表名
	CompareField string `mapstructure:"compareField" json:"compareField" yaml:"compareField"` // 需要比较时间的字段
	Interval     string `mapstructure:"interval" json:"interval" yaml:"interval"`             // 时间间隔
}

// Config 整体配置
type Config struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Timer  Timer  `mapstructure:"timer" json:"timer" yaml:"timer"`
}

// Redis 前缀配置
var (
	MenueIds = "Menu_"
)
