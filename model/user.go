package model

import (
	"time"
	"net/http"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (f *User) Serialize() *UserSerializer {
	return &UserSerializer {
		ID: f.ID,
		Username: f.Username,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}

type UserSerializer struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (f *UserSerializer) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type CurrentCharacter struct {
	CharacterID *uint `json:"character_id"`
}

func (f *CurrentCharacter) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
