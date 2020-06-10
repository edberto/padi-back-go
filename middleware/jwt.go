package middleware

import (
	"context"
	"fmt"
	"net/http"
	"padi-back-go/config"
	"padi-back-go/helper"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userTokenCollection = "user_tokens"

type AccessDetails struct {
	UUID      string `bson:"uuid"`
	UserID    string `bson:"user_id"`
	ExpiredAt string `bson:"expired_at"`
}

func SetTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ad := new(AccessDetails)

		req := c.Request

		cfg := config.NewConfig("config.yaml")
		key := cfg.GetString("key.access")

		jwt := helper.NewJWT(key)
		claims, err := jwt.ExtractClaims(req)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid Token!"))
			return
		}

		db := GetMongoDB(c)
		col := db.Collection(userTokenCollection)
		err = col.FindOne(req.Context(), bson.D{{"uuid", claims["access-uuid"]}}).Decode(&ad)
		if err != nil && err == mongo.ErrNoDocuments {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Token not found!"))
			return
		}
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Internal Server Error!"))
			return
		}

		userID, _ := strconv.Atoi(ad.UserID)
		c.Request = req.WithContext(context.WithValue(req.Context(), "user_id", userID))
		c.Request = c.Request.WithContext(context.WithValue(req.Context(), "access-uuid", claims["access-uuid"]))

		c.Next()
	}
}
