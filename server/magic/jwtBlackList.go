package magic

import (
	"time"

	"github.com/jianyuezhexue/MagicAdmin/model/system"
	"go.uber.org/zap"
)

// JwtService struct
type JwtService struct {
}

// JSONInBlacklist 拉黑jwt
func (jwtService *JwtService) JSONInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = Orm.Create(&jwtList).Error
	if err != nil {
		return
	}
	LocalCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

// IsBlacklist 判断JWT是否在黑名单内部
func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := LocalCache.Get(jwt)
	return ok
	//err := Orm.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	//isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	//return !isNotFound
}

// GetRedisJWT 获取Redis中的JWT
func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = Redis.Get(userName)
	return redisJWT, err
}

// SetRedisJWT jwt存入redis并设置过期时间
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(Config.JWT.ExpiresTime) * time.Second
	err = Redis.Set(userName, jwt, timer)
	return err
}

// LoadAll 加载全部黑名单到缓存中
func LoadAll() {
	var data []string
	err := Orm.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		Logger.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		LocalCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
