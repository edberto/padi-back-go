package register

import (
	"log"
	"net/http"
	"padi-back-go/helper"

	"github.com/gin-gonic/gin"
)

var (
	registerUser = Handler.RegisterUser
)

type IHandler interface {
	RegisterHandler(c *gin.Context)
}

type Handler struct {
	IUsecase
}

func NewHandler(u IUsecase) IHandler {
	return &Handler{u}
}

func (h Handler) RegisterHandler(c *gin.Context) {
	req := new(RegisterR)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Wrap(nil, "Bad Request!"))
		return
	}

	registerP := new(RegisterUserP)
	(*registerP).Username, (*registerP).Password = (*req).Username, (*req).Password

	res, err := registerUser(h, c, registerP)
	if err != nil && err == helper.ErrUserExisted {
		c.JSON(http.StatusBadRequest, helper.Wrap(nil, "User Existed!"))
		log.Printf("Error: %v", err)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Wrap(nil, "Internal Server Error!"))
		log.Printf("Error Occured: %v", err)
		return
	}

	c.JSON(http.StatusOK, helper.Wrap(*res, "Success"))
	return
}
