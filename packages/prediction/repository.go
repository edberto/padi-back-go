package prediction

import (
	"fmt"
	"padi-back-go/helper"
	"padi-back-go/middleware"

	"github.com/gin-gonic/gin"
)

type IRepository interface {
	InsertOne(c *gin.Context, p *InsertOneParam) (res *UserPrediction, err error)
	FindAll(c *gin.Context, p *FindAllParam) (res *[]UserPrediction, err error)
}

type Repository struct{}

func NewRepository() IRepository {
	return &Repository{}
}

type InsertOneParam struct {
	ImagePath  string
	Prediction int
	UserID     int
}

func (r *Repository) InsertOne(c *gin.Context, p *InsertOneParam) (res *UserPrediction, err error) {
	res = new(UserPrediction)

	q := `
		INSERT INTO public.predictions (user_id, prediction, image_path) 
		VALUES ($1, $2, $3)
		RETURNING user_id, prediction, image_path`

	db := middleware.GetPostgres(c)
	err = db.QueryRow(q, p.UserID, p.Prediction, p.ImagePath).Scan(&res.UserID, &res.Prediction, &res.ImagePath)

	return res, err
}

type FindAllParam struct {
	UserID int
}

func (r *Repository) FindAll(c *gin.Context, p *FindAllParam) (res *[]UserPrediction, err error) {
	res = new([]UserPrediction)

	selectQ := `
		SELECT prediction, user_id, image_path, updated_at
		FROM public.predictions
	`
	conditionQ := ` WHERE deleted_at IS NULL`
	conditionP := new([]interface{})
	if a := p.UserID; a != 0 {
		conditionQ += ` AND user_id = ?`
		*conditionP = append(*conditionP, p.UserID)
	}

	limitQ := ` ORDER BY updated_at DESC`

	q := fmt.Sprint(selectQ, conditionQ, limitQ)
	q = helper.ReplacePlaceholder(q, 1)

	db := middleware.GetPostgres(c)
	rows, err := db.Query(q, *conditionP...)
	if err != nil {
		return res, err
	}

	defer rows.Close()
	for rows.Next() {
		t := new(UserPrediction)

		if err := rows.Scan(&t.Prediction, &t.UserID, &t.ImagePath, &t.UpdatedAt); err != nil {
			return res, err
		}

		*res = append(*res, *t)
	}

	return res, err
}
