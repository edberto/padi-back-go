package prediction

import (
	"net/http"
	"padi-back-go/helper"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	FindAllHandler(c *gin.Context)
	InsertOneHandler(c *gin.Context)
}

type Handler struct {
	IUsecase
}

func NewHandler(u IUsecase) IHandler {
	return &Handler{u}
}

func (h *Handler) FindAllHandler(c *gin.Context) {
	findAllParam := new(FindAllUCParam)
	findAllParam.UserID = c.Request.Context().Value("user_id").(int)
	res, err := h.GetAll(c, findAllParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Wrap(nil, "Internal Server Error!"))
		return
	}

	c.JSON(http.StatusOK, helper.Wrap(*res, "Success"))

}

func (h *Handler) InsertOneHandler(c *gin.Context) {
	req := new(PredictionR)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Wrap(nil, "Bad Request!"))
		return
	}

	insertOneP := new(InsertOneUCParam)
	insertOneP.ImagePath, insertOneP.Prediction, insertOneP.UserID = req.ImagePath, req.Prediction, c.Request.Context().Value("user_id").(int)
	res, err := h.StoreOne(c, insertOneP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Wrap(nil, "Internal Server Error!"))
		return
	}

	c.JSON(http.StatusOK, helper.Wrap(*res, "Success"))
}
