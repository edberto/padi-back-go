package setup

import (
	"context"
	"log"
	"padi-back-go/config"
	"padi-back-go/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func setupMongoDB(r *gin.Engine, cfg config.IConfig) {
	uri := cfg.GetString("database.mongo.uri")

	client, err := middleware.ConnectMongoDB(uri)
	if err != nil {
		log.Fatalf("Error when connecting to db: %v", err)
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalf("Failed to ping mongo db client: %v", err)
	} else {
		log.Println("Mongo DB Client is successfully connected")
	}

	dbName := cfg.GetString("database.mongo.db-name")
	db := middleware.FindDatabase(client, dbName)

	mw := middleware.SetMongoDBCtx(db)
	r.Use(mw)
}
