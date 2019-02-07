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

type CharacterSerializer struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
}

func (f *CharacterSerializer) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (self *Character) Serialize() *CharacterSerializer {
	return &CharacterSerializer{
		ID: self.ID,
		Name: self.Name,
	}
}
