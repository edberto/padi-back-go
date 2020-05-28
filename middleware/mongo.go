package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoDBA = "mongo-db"
)

func ConnectMongoDB(uri string) (*mongo.Client, error) {
	c, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return c, err
	}

	err = c.Connect(context.Background())

	return c, err
}

func FindDatabase(c *mongo.Client, dbName string) *mongo.Database {
	return c.Database(dbName)
}

func SetMongoDBCtx(db *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(mongoDBA, db)
		c.Next()
	}
}

func GetMongoDB(c *gin.Context) *mongo.Client {
	return c.MustGet(mongoDBA).(*mongo.Client)
}
