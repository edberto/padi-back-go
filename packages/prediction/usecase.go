package prediction

import "github.com/gin-gonic/gin"

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
		(*t).ImagePath, (*t).Prediction, (*t).UserID, (*t).UpdatedAt = p.ImagePath, p.Prediction, p.UserID, p.UpdatedAt

		*res = append(*res, *t)
	}

	return res, err
}

type InsertOneUCParam struct {
}

func (u Usecase) StoreOne(c *gin.Context, p *InsertOneUCParam) (res *PredictionVM, err error) {
	return res, err
}
