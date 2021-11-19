package magic

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jianyuezhexue/MagicAdmin/model/system/request"
)

// JWT jwt结构体
type JWT struct {
	SigningKey []byte
}

var (
	// ErrTokenExpired 存在
	ErrTokenExpired = errors.New("Token is expired")
	// ErrTokenNotValidYet 未验证
	ErrTokenNotValidYet = errors.New("Token not active yet")
	// ErrTokenMalformed 格式错误
	ErrTokenMalformed = errors.New("That's not even a token")
	// ErrTokenInvalid 无法验证
	ErrTokenInvalid = errors.New("Couldn't handle this token")
)

// NewJWT 加入签名实例化JWT
func NewJWT() *JWT {
	return &JWT{
		[]byte(Config.JWT.SigningKey),
	}
}

// CreateClaims 创建Claims
func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: Config.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                   // 签名生效时间
			ExpiresAt: time.Now().Unix() + Config.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    Config.JWT.Issuer,                          // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := SingleFlight.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid

	}
	return nil, ErrTokenInvalid
}
