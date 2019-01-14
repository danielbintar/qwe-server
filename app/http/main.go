package main

import (
	"fmt"
	"time"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/danielbintar/go-record/db"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Post("/users/sign_in", Login)

	fmt.Println("listen to 3333")
	http.ListenAndServe(":3333", r)
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

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

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}
