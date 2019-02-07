package model

import (
	"time"
	"net/http"
)

type Character struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (f *Character) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
