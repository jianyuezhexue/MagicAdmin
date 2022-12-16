package middle

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
)

var jwtService = magic.JwtService{}

// JWTAuth jwt鉴权
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token字符串
		token := c.Request.Header.Get("x-token")
		if token == "" {
			magic.HttpFail(c, http.StatusForbidden, "未登录或非法访问", "")
			c.Abort()
			return
		}

		// 校验授权
		jwt := magic.NewJWT()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == magic.ErrTokenExpired {
				magic.HttpFail(c, http.StatusForbidden, "授权已过期", "")
				c.Abort()
				return
			}
			magic.HttpFail(c, http.StatusForbidden, err.Error(), "")
			c.Abort()
			return
		}

		// 过期自动续期
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + magic.Config.JWT.ExpiresTime
			newToken, _ := jwt.CreateTokenByOldToken(token, *claims)
			newClaims, _ := jwt.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if magic.Config.System.UseMultiPoint {
				// 重置新Token
				_ = jwtService.SetRedisJWT(newToken, newClaims.UserName)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}
