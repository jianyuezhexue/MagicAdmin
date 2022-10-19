package middle

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	"go.uber.org/zap"
)

var jwtService = magic.JwtService{}

// JWTAuth jwt鉴权
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token字符串
		token := c.Request.Header.Get("x-token")
		if token == "" {
			magic.Fail(c, http.StatusForbidden, "未登录或非法访问", "")
			c.Abort()
			return
		}

		// 校验授权
		jwt := magic.NewJWT()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == magic.ErrTokenExpired {
				magic.Fail(c, http.StatusForbidden, "授权已过期", "")
				c.Abort()
				return
			}
			magic.Fail(c, http.StatusForbidden, err.Error(), "")
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
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.UserName)
				if err != nil {
					magic.Logger.Error("get redis jwt failed", zap.Any("err", err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JSONInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.UserName)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}
