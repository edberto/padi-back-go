package register

import (
	"fmt"
	"padi-back-go/helper"
	"padi-back-go/middleware"

	"github.com/gin-gonic/gin"
)

type IRepo interface {
	InsertOne(c *gin.Context, p *InsertOneParam) (res *Model, err error)
	FindOne(c *gin.Context, p *FindOneParam) (res *Model, err error)
}

type Repo struct{}

func NewRepo() IRepo {
	return &Repo{}
}

type InsertOneParam struct {
	Username string
	Password string
}

func (r *Repo) InsertOne(c *gin.Context, p *InsertOneParam) (res *Model, err error) {
	res = new(Model)

	q := `
		INSERT INTO public.users (username, password) VALUES (?, ?)
		RETURNING id, username, password
	`
	q = helper.ReplacePlaceholder(q, 1)

	db := middleware.GetPostgres(c)
	err = db.QueryRow(q, (*p).Username, (*p).Password).Scan(&res.ID, &res.Username, &res.Password)
	if err != nil {
		return res, err
	}

	return res, err
}

type FindOneParam struct {
	Username string
}

func (r *Repo) FindOne(c *gin.Context, p *FindOneParam) (res *Model, err error) {
	res = new(Model)

	selectQ := `
		SELECT id, username, password
		FROM public.users
	`

	conditionQ := ` WHERE deleted_at IS NULL`
	conditionP := new([]interface{})
	if a := &(*p).Username; *a != "" {
		conditionQ += ` AND username = ?`
		*conditionP = append(*conditionP, *a)
	}

	limitQ := ` ORDER BY updated_at DESC LIMIT 1`

	q := fmt.Sprint(selectQ, conditionQ, limitQ)
	q = helper.ReplacePlaceholder(q, 1)

	db := middleware.GetPostgres(c)
	err = db.QueryRow(q, *conditionP...).Scan(&res.ID, &res.Username, &res.Password)
	if err != nil {
		return nil, err
	}

	return res, nil
}
