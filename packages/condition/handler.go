package condition

import (
	"net/http"
	"padi-back-go/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	FindLabelHandler(c *gin.Context) 
}

type Handler struct {
	IUsecase
}

func NewHandler(uc IUsecase) IHandler {
	return &Handler{uc}
}

func (h Handler) FindLabelHandler(c *gin.Context) {
	labelID, err := strconv.Atoi(c.Param("label-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.Wrap(nil, "Bad Request"))
		return
	}

	vm, err := h.FindConditionDetail(c, labelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Wrap(nil, "Internal Server Error"))
		return
	}
	
	c.JSON(http.StatusOK, helper.Wrap(vm, "Success"))
	return
}