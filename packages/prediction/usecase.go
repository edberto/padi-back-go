package prediction

import (
	"padi-back-go/helper"
	"time"

	"github.com/gin-gonic/gin"
)

type IUsecase interface {
	GetAll(c *gin.Context, p *FindAllUCParam) (res *[]PredictionVM, err error)
	StoreOne(c *gin.Context, p *InsertOneUCParam) (res *PredictionVM, err error)
}

type Usecase struct {
	IRepository
}

func NewUsecase(r IRepository) IUsecase {
	return &Usecase{r}
}

type FindAllUCParam struct {
	UserID int
}

func (u Usecase) GetAll(c *gin.Context, p *FindAllUCParam) (res *[]PredictionVM, err error) {
	res = new([]PredictionVM)

	findAllP := new(FindAllParam)
	findAllP.UserID = p.UserID

	predictions, err := u.FindAll(c, findAllP)
	if err != nil {
		return res, err
	}

	for _, p := range *predictions {
		t := new(PredictionVM)
		(*t).ImagePath, (*t).Prediction, (*t).UserID, (*t).UpdatedAt, (*t).Label = p.ImagePath, p.Prediction, p.UserID, p.UpdatedAt.In(helper.WIB), helper.LabelItoa[p.Prediction]

		*res = append(*res, *t)
	}

	return res, err
}

type InsertOneUCParam struct {
	ImagePath  string
	Prediction int
	UserID     int
}

func (u Usecase) StoreOne(c *gin.Context, p *InsertOneUCParam) (res *PredictionVM, err error) {
	res = new(PredictionVM)

	insertOneP := new(InsertOneParam)
	insertOneP.ImagePath, insertOneP.Prediction, insertOneP.UserID = p.ImagePath, p.Prediction, p.UserID

	prediction, err := u.InsertOne(c, insertOneP)
	if err != nil {
		return res, err
	}

	res.ImagePath, res.Prediction, res.UserID, res.Label, res.UpdatedAt = prediction.ImagePath, prediction.Prediction, prediction.UserID, helper.LabelItoa[prediction.Prediction], time.Now().In(helper.WIB)

	return res, err
}
