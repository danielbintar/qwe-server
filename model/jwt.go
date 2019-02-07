package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	UserId uint
	jwt.StandardClaims
}

func NewJwt(user *User) *Jwt {
	return &Jwt{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 7200,
		},
	}
}
