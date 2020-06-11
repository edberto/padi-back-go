package setup

import (
	"padi-back-go/config"
	"padi-back-go/middleware"

	"github.com/gin-gonic/gin"
)

func setupConfig(r *gin.Engine, cfg config.IConfig) {
	mw := middleware.SetCfgCtx(cfg)
	r.Use(mw)
}
