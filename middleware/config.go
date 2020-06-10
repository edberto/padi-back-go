package middleware

import (
	"padi-back-go/config"

	"github.com/gin-gonic/gin"
)

var (
	cfgA = "config"
)

func SetCfgCtx(cfg config.IConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(cfgA, cfg)
		c.Next()
	}
}

func GetConfig(c *gin.Context) config.IConfig {
	return c.MustGet(cfgA).(config.IConfig)
}
