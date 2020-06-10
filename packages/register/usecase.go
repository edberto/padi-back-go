package register

import (
	"database/sql"
	"padi-back-go/helper"

	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"
	"golang.org/x/crypto/bcrypt"
)

var (
	findOne   = Usecase.FindOne
	insertOne = Usecase.InsertOne
)

type IUsecase interface {
	RegisterUser(c *gin.Context, p *RegisterUserP) (res *RegisterVM, err error)
}

type Usecase struct {
	IRepo
}

func NewUsecase(r IRepo) IUsecase {
	return &Usecase{r}
}

type RegisterUserP struct {
	Username string
	Password string
}

func (u *Usecase) RegisterUser(c *gin.Context, p *RegisterUserP) (res *RegisterVM, err error) {
	res = new(RegisterVM)

	findOneP := new(FindOneParam)
	(*findOneP).Username = (*p).Username

	user, err := findOne(*u, c, findOneP)
	if err != nil && err != sql.ErrNoRows {
		return res, stacktrace.Propagate(err, "Failed to find user")
	}
	if user != nil {
		return res, helper.ErrUserExisted
	}

	pwd := []byte((*p).Password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return res, stacktrace.Propagate(err, "Failed to encrypt password")
	}

	insertOneP := new(InsertOneParam)
	(*insertOneP).Username, (*insertOneP).Password = (*p).Username, string(hash)

	user, err = insertOne(*u, c, insertOneP)
	if err != nil {
		return res, stacktrace.Propagate(err, "Failed to insert user")
	}

	(*res).ID, (*res).Username, (*res).Password = (*user).ID, (*user).Username, (*user).Password

	return res, err
}
