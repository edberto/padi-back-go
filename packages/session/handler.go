package session

import (
	"net/http"
	"padi-back-go/config"
	"padi-back-go/helper"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type IHandler interface {
	LoginHandler(c *gin.Context)
	RefreshHandler(c *gin.Context)
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
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Wrap(nil, "Bad Request!"))
	}

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

func (h *Handler) RefreshHandler(c *gin.Context) {
	req := new(RefreshR)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Wrap(nil, "Bad Request!"))
	}

	cfg := config.NewConfig("config.yaml")
	key := cfg.GetString("key.refresh")

	j := helper.NewJWT(key)
	token, err := j.VerifyToken((*req).Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.Wrap(nil, "Invalid Token!"))
		return
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusBadRequest, helper.Wrap(nil, "Invalid Token!"))
		return
	}

	refreshP := new(RefreshParam)
	(*refreshP).UUID = claim["refresh-uuid"].(string)
	res, err := h.Refresh(c, refreshP)
	if err != nil && err == helper.ErrTokenExpired {
		c.JSON(http.StatusBadRequest, helper.Wrap(nil, "Please Login"))
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Wrap(nil, "Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, helper.Wrap(*res, "Success"))
}

func (h *Handler) LogoutHandler(c *gin.Context) {
	uuid := c.Request.Context().Value("access-uuid").(string)
	logoutP := new(LogoutParam)
	(*logoutP).UUID = uuid

	err := h.Logout(c, logoutP)
	if err != nil && err == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, helper.Wrap(nil, "Token Expired!"))
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Wrap(nil, "Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, helper.Wrap(nil, "Success"))
}
