package condition

import (
	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"
)

type IUsecase interface {
	FindConditionDetail(c *gin.Context, labelID int) (vm ConditionDetailVM, err error)
}

type Usecase struct {
	IRepository
}

func NewUsecase(rp IRepository) IUsecase {
	return &Usecase{rp}
}

func (uc Usecase) FindConditionDetail(c *gin.Context, labelID int) (vm ConditionDetailVM, err error) {
	info, err := uc.FindOneByLabelID(c, labelID)
	if err != nil {
		return vm, stacktrace.Propagate(err, "Failed to find condition")
	}

	vm = ConditionDetailVM{
		Label: info.Label,
		Description: info.Desciption,
		Effect: info.Effect,
		Solution: info.Solution,
		Prevention: info.Prevention,
	}
	return vm, err
}