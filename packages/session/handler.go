package session

import (
	"net/http"
	"padi-back-go/helper"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	LoginHandler(c *gin.Context)
	LogoutHandler(c *gin.Context)
}

type Handler struct {
	IUsecase
}

func NewHandler(u IUsecase) IHandler {
	return &Handler{u}
}

func (h *Handler) LoginHandler(c *gin.Context) {
	req := new(LoginR)
	c.ShouldBindJSON(&req)

	loginP := new(LoginParam)
	loginP.Username, loginP.Password = req.Username, req.Password
	res, err := h.Login(c, loginP)
	if err != nil && err == helper.ErrUserNotFound {
		c.JSON(http.StatusBadRequest, helper.Wrap(nil, "User not found!"))
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Wrap(nil, "Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, helper.Wrap(*res, "Success"))
}

func (h *Handler) LogoutHandler(c *gin.Context) {

}
