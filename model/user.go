package model

import (
	"time"
	"net/http"
)

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (f *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
