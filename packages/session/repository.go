package session

import (
	"fmt"
	"padi-back-go/helper"
	"padi-back-go/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userTokenCollection = "user_tokens"

type IRepository interface {
	FindOneUser(c *gin.Context, p *FindOneUserParam) (res *User, err error)
	StoreOneToken(c *gin.Context, p *StoreOneTokenParam) (res primitive.ObjectID, err error)
	DeleteOneToken(c *gin.Context, p *DeleteOneTokenParam) (err error)
}

type Repository struct{}

func NewRepository() IRepository {
	return &Repository{}
}

type FindOneUserParam struct {
	Username string
	Password string
}

func (r *Repository) FindOneUser(c *gin.Context, p *FindOneUserParam) (res *User, err error) {
	res = new(User)

	selectQ := `
		SELECT id, username, password
		FROM public.users
	`

	conditionQ := ` WHERE deleted_at IS NULL`
	conditionP := new([]interface{})
	if a := (*p).Username; a != "" {
		conditionQ += ` AND username = ?`
		*conditionP = append(*conditionP, a)
	}
	if a := (*p).Password; a != "" {
		conditionQ += ` AND password = ?`
		*conditionP = append(*conditionP, a)
	}

	limitQ := ` ORDER BY updated_at DESC LIMIT 1`

	q := fmt.Sprint(selectQ, conditionQ, limitQ)
	q = helper.ReplacePlaceholder(q, 1)

	db := middleware.GetPostgres(c)
	err = db.QueryRow(q, *conditionP...).Scan(&res.ID, &res.Username, &res.Password)

	return res, err
}

type StoreOneTokenParam struct {
	UUID      string
	UserID    int
	ExpiredAt time.Time
}

func (r *Repository) StoreOneToken(c *gin.Context, p *StoreOneTokenParam) (res primitive.ObjectID, err error) {
	db := middleware.GetMongoDB(c)
	col := db.Collection(userTokenCollection)

	result, err := col.InsertOne(c.Request.Context(), map[string]string{
		"uuid":       (*p).UUID,
		"user_id":    fmt.Sprint((*p).UserID),
		"expired_at": (*p).ExpiredAt.String(),
	})

	return result.InsertedID.(primitive.ObjectID), err
}

type DeleteOneTokenParam struct {
}

func (r *Repository) DeleteOneToken(c *gin.Context, p *DeleteOneTokenParam) (err error) {
	return err
}
