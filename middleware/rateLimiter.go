package middlewares

import (
	"beianAPI/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func RateLimiter(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)

	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.JSON(http.StatusOK, gin.H{
				"code":    errmsg.ERROR_RATE_LIMIT,
				"message": errmsg.GetErrMsg(errmsg.ERROR_RATE_LIMIT),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
