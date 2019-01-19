package controller

import (	
	"time"
	"net/http"

	"github.com/danielbintar/go-record/db"
	"github.com/go-chi/render"
)

func Login(w http.ResponseWriter, r *http.Request) {
	data := &User{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if data.Valid() {
		render.Render(w, r, data)
	} else {
		render.Render(w, r, ErrNotFound)
	}
}

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

func (f *User) Bind(r *http.Request) error {
	return nil
}

func (f *User) Valid() bool {
	err := db.FindBy(&f, []string{"username", "=", f.Username}, []string{"password", "=", f.Password});
	return err == nil
}
