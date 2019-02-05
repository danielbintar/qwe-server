package model

import (
	"net/http"
	"time"

	"github.com/danielbintar/qwe-server/lib"
)

type Token struct {
	Token     string     `json:"token"`
	ExpiresAt time.Time  `json:"expires_at"`
}

func (f *Token) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewToken(token string, jwt *Jwt) *Token {
	return &Token{
		Token: token,
		ExpiresAt: lib.ParseUnix(jwt.ExpiresAt),
	}
}
