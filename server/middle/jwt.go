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
		token := c.Request.Header.Get("x-token")
		if token == "" {
			magic.Fail(c, http.StatusForbidden, "未登录或非法访问")
			c.Abort()
			return
		}
		jwt := magic.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == magic.ErrTokenExpired {
				magic.Fail(c, http.StatusForbidden, "授权已过期")
				c.Abort()
				return
			}
			magic.Fail(c, http.StatusForbidden, err.Error())
			c.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + magic.Config.JWT.ExpiresTime
			newToken, _ := jwt.CreateTokenByOldToken(token, *claims)
			newClaims, _ := jwt.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if magic.Config.System.UseMultipoint {
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
