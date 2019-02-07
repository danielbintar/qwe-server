package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	UserID uint
	jwt.StandardClaims
}

func NewJwt(user *User) *Jwt {
	return &Jwt{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 7200,
		},
	}
}
