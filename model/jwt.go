package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	UserID   uint
	Username string
	jwt.StandardClaims
}

func NewJwt(user *User) *Jwt {
	return &Jwt{
		UserID: user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 7200,
		},
	}
}
