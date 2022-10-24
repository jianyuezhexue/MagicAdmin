package middle

import (
	"time"

	"github.com/gin-gonic/gin"
)

// 统计耗时中间件
func CostTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UnixNano()
		c.Set("Magic_sTime", start)
		c.Next()
	}
}
