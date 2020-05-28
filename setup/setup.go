package setup

import (
	"padi-back-go/config"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	
	cfg := config.NewConfig("config.yaml")
	setupMongoDB(r, cfg)

	return r
}
