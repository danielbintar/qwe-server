package model

import (
	"time"
	"net/http"

	"github.com/go-chi/render"
)

type Character struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// not persisted
	ActivePlace *string `json:"active_place" gorm:"-"`
}

type CharacterSerializer struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	ActivePlace *string `json:"active_place"`
}

func (f *CharacterSerializer) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (self *Character) Serialize() render.Renderer {
	return &CharacterSerializer{
		ID: self.ID,
		Name: self.Name,
		ActivePlace: self.ActivePlace,
	}
}

type CharacterPosition struct {
	ID uint `json:"id"`
	X  uint `json:"x"`
	Y  uint `json:"y"`
}

type CharacterLogout struct {
	ID uint `json:"id"`
}
