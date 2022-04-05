package routes

import (
	v1 "beianAPI/api/v1"
	middlewares "beianAPI/middleware"
	"beianAPI/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"time"
)

const (
	rateLimitInternal = time.Minute / 12
	rateLimitCap      = 10
)

func InitRouter() {
	gin.SetMode(utils.Conf.Mode)
	r := gin.New()
	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.RateLimiter(rateLimitInternal, rateLimitCap))

	apiGroup := r.Group("/api/v1")
	user := apiGroup.Group("user")
	{
		user.POST("signUp", v1.SignUp)
		user.POST("getToken", v1.GetToken)
	}

	beianInfo := apiGroup.Group("beianInfo")
	beianInfo.Use(middlewares.JWY())
	{
		beianInfo.GET("", v1.GetInfo)
		beianInfo.GET("latest", v1.GetLatestInfo)
	}

	_ = r.Run(utils.Conf.HTTPPort)
}
