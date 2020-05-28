package condition

import (
	"padi-back-go/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	conditionCollection = "plant_conditions"
)

type IRepository interface {
	FindOneByLabelID(c *gin.Context, labelID int) (res ConditionModel, err error)
}

type Repository struct{}

func NewRepository() IRepository {
	return &Repository{}
}

func (rp Repository) FindOneByLabelID(c *gin.Context, labelID int) (res ConditionModel, err error) {
	db := middleware.GetMongoDB(c)
	col := db.Collection(conditionCollection)
	err = col.FindOne(c.Request.Context(), bson.D{{"label_id", labelID}}).Decode(&res)
	return res, err
}
