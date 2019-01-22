package model

import (
	"time"
	"net/http"
)

type User struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserPosition struct {
	Id uint `json:"id"`
	X  uint `json:"x"`
	Y  uint `json:"y"`
}

func (f *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
