package magic

// JwtService struct
type JwtService struct {
}

// GetRedisJWT 获取Redis中的JWT
func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = Redis.Get(userName)
	return redisJWT, err
}

// SetRedisJWT jwt存入redis并设置过期时间
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	// timer := time.Duration(Config.JWT.ExpiresTime) * time.Second
	timer := Config.JWT.ExpiresTime
	err = Redis.Set(userName, jwt, timer)
	return err
}
