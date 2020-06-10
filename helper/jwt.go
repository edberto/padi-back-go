package helper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type IJWT interface {
	AddClaim(key string, value interface{}) *JWT
	CreateToken() (token string, err error)
	ExtractClaims(r *http.Request) (res map[string]interface{}, err error)
}

type JWT struct {
	IJWT
	Key   string
	Claim map[string]interface{}
}

func NewJWT(key string) *JWT {
	claim := map[string]interface{}{}
	return &JWT{Key: key, Claim: claim}
}

func (j *JWT) AddClaim(key string, value interface{}) *JWT {
	claim := j.Claim
	claim[key] = value
	j.Claim = claim
	return j
}

func (j *JWT) CreateToken() (token string, err error) {
	claim := jwt.MapClaims{}
	for k, v := range j.Claim {
		claim[k] = v
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return t.SignedString([]byte(j.Key))
}

func (j *JWT) ExtractClaims(r *http.Request) (res map[string]interface{}, err error) {
	res = map[string]interface{}{}

	tokenA := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	token, err := j.verifyToken(tokenA)
	if err != nil {
		return res, err
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return res, fmt.Errorf("Invalid Token")
	}
	return claim, err
}

func (j *JWT) verifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Wrong signing method: %v", token.Header["alg"])
		}
		return []byte(j.Key), nil
	})
}
