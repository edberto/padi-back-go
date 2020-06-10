package setup

import (
	"log"
	"os"
	"padi-back-go/config"
	"padi-back-go/middleware"

	"github.com/gin-gonic/gin"
)

func setupPostgresql(r *gin.Engine, cfg config.IConfig) {
	uri := os.Getenv("DATABASE_URL")
	if uri == "" {
		uri = cfg.GetString("database.postgres.uri")
	}

	db, err := middleware.ConnectPostgresql(uri)
	if err != nil {
		log.Fatalf("Failed to initiate postgresql: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping postgresql: %v", err)
	} else {
		log.Println("Postgresql DB is successfully connected")
	}

	mw := middleware.SetPostgresCtx(db)
	r.Use(mw)
}