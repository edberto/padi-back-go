package session

import (
	"database/sql"
	"padi-back-go/config"
	"padi-back-go/helper"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/palantir/stacktrace"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type IUsecase interface {
	Login(c *gin.Context, p *LoginParam) (res *LoginVM, err error)
	Refresh(c *gin.Context, p *RefreshParam) (res *LoginVM, err error)
	Logout(c *gin.Context, p *LogoutParam) (err error)
}

type Usecase struct {
	IRepository
}

func NewUsecase(r IRepository) IUsecase {
	return &Usecase{r}
}

type LoginParam struct {
	Username string
	Password string
}

func (u *Usecase) Login(c *gin.Context, p *LoginParam) (res *LoginVM, err error) {
	res = new(LoginVM)

	findOneUserP := new(FindOneUserParam)
	(*findOneUserP).Username = (*p).Username

	user, err := u.FindOneUser(c, findOneUserP)
	if err != nil && err == sql.ErrNoRows {
		return res, helper.ErrUserNotFound
	}
	if err != nil {
		return res, stacktrace.Propagate(err, "Failed to get user data")
	}

	actualPassword := []byte((*user).Password)
	claimedPassword := []byte((*p).Password)
	if err := bcrypt.CompareHashAndPassword(actualPassword, claimedPassword); err != nil {
		return res, helper.ErrUserNotFound
	}

	config := config.NewConfig("config.yaml")
	accessKey := config.GetString("key.access")
	refreshKey := config.GetString("key.refresh")

	accessJWT := helper.NewJWT(accessKey)
	accessExpire := time.Now().Add(3 * 24 * time.Hour)
	accessUUID := uuid.New().String()
	accessTokenString, err := accessJWT.AddClaim("user_id", (*user).ID).
		AddClaim("expired_at", accessExpire).
		AddClaim("access-uuid", accessUUID).
		CreateToken()
	if err != nil {
		return res, stacktrace.Propagate(err, "Error when generating access token")
	}

	refreshJWT := helper.NewJWT(refreshKey)
	refreshExpire := time.Now().Add(7 * 24 * time.Hour)
	refreshUUID := uuid.New().String()
	refreshTokenString, err := refreshJWT.AddClaim("user_id", (*user).ID).
		AddClaim("expired_at", refreshExpire).
		AddClaim("refresh-uuid", refreshUUID).
		CreateToken()
	if err != nil {
		return res, stacktrace.Propagate(err, "Error when generating refresh token")
	}

	storeAccessTokenP := new(StoreOneTokenParam)
	(*storeAccessTokenP).UserID, (*storeAccessTokenP).UUID, (*storeAccessTokenP).ExpiredAt = (*user).ID, accessUUID, accessExpire
	_, err = u.StoreOneToken(c, storeAccessTokenP)
	if err != nil {
		return res, stacktrace.Propagate(err, "Error when storing token")
	}

	storeRefreshTokenP := new(StoreOneTokenParam)
	(*storeRefreshTokenP).UserID, (*storeRefreshTokenP).UUID, (*storeRefreshTokenP).ExpiredAt = (*user).ID, refreshUUID, refreshExpire
	_, err = u.StoreOneToken(c, storeRefreshTokenP)
	if err != nil {
		return res, stacktrace.Propagate(err, "Error when storing token")
	}

	(*res).AccessToken, (*res).RefreshToken = accessTokenString, refreshTokenString
	return res, err
}

type RefreshParam struct {
	UUID string
}

func (u *Usecase) Refresh(c *gin.Context, p *RefreshParam) (res *LoginVM, err error) {
	res = new(LoginVM)

	findOneTokenP := new(FindOneTokenParam)
	(*findOneTokenP).UUID = (*p).UUID
	token, err := u.FindOneToken(c, findOneTokenP)
	if err != nil && err == mongo.ErrNoDocuments {
		return res, helper.ErrTokenExpired
	}
	if err != nil {
		return res, stacktrace.Propagate(err, "Unable to fetch token")
	}

	config := config.NewConfig("config.yaml")
	accessKey := config.GetString("key.access")
	refreshKey := config.GetString("key.refresh")

	accessJWT := helper.NewJWT(accessKey)
	accessExpire := time.Now().Add(3 * 24 * time.Hour)
	accessUUID := uuid.New().String()
	accessTokenString, err := accessJWT.AddClaim("user_id", token.UserID).
		AddClaim("expired_at", accessExpire).
		AddClaim("access-uuid", accessUUID).
		CreateToken()
	if err != nil {
		return res, stacktrace.Propagate(err, "Error when generating access token")
	}

	refreshJWT := helper.NewJWT(refreshKey)
	refreshExpire := time.Now().Add(7 * 24 * time.Hour)
	refreshUUID := uuid.New().String()
	refreshTokenString, err := refreshJWT.AddClaim("user_id", token.UserID).
		AddClaim("expired_at", refreshExpire).
		AddClaim("refresh-uuid", refreshUUID).
		CreateToken()
	if err != nil {
		return res, stacktrace.Propagate(err, "Error when generating refresh token")
	}

	userIDI, _ := strconv.Atoi(token.UserID)
	storeAccessTokenP := new(StoreOneTokenParam)
	(*storeAccessTokenP).UserID, (*storeAccessTokenP).UUID, (*storeAccessTokenP).ExpiredAt = userIDI, accessUUID, accessExpire
	_, err = u.StoreOneToken(c, storeAccessTokenP)
	if err != nil {
		return res, stacktrace.Propagate(err, "Error when storing token")
	}

	storeRefreshTokenP := new(StoreOneTokenParam)
	(*storeRefreshTokenP).UserID, (*storeRefreshTokenP).UUID, (*storeRefreshTokenP).ExpiredAt = userIDI, refreshUUID, refreshExpire
	_, err = u.StoreOneToken(c, storeRefreshTokenP)
	if err != nil {
		return res, stacktrace.Propagate(err, "Error when storing token")
	}

	(*res).AccessToken, (*res).RefreshToken = accessTokenString, refreshTokenString
	return res, err
}

type LogoutParam struct {
	UUID string
}

func (u *Usecase) Logout(c *gin.Context, p *LogoutParam) (err error) {
	deleteOneTokenP := new(DeleteOneTokenParam)
	(*deleteOneTokenP).UUID = (*p).UUID

	err = u.DeleteOneToken(c, deleteOneTokenP)

	return err
}
