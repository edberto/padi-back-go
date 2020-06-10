package setup

import (
	"padi-back-go/config"
	"padi-back-go/route"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	cfg := config.NewConfig("config.yaml")
	setupConfig(r, cfg)
	setupMongoDB(r, cfg)
	setupPostgresql(r, cfg)

	route.Initialize(r)
	return r
}
