package controller

import (
	"os"
	"net/http"
	"encoding/json"

	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/service/auth"

	"github.com/go-chi/render"

	jwt "github.com/dgrijalva/jwt-go"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var form auth.LoginForm

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}


	userI, errors := auth.Login(form)
	if errors == nil {
		byteData, _ := json.Marshal(userI)
		var user *model.User
		json.Unmarshal(byteData, &user)

		token := createToken(user)
		render.Render(w, r, token)
	} else {
		render.Render(w, r, ErrNotFound)
	}
}

func createToken(user *model.User) *model.Token {
	tk := model.NewJwt(user)
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("APPLICATION_SECRET_KEY")))

	return model.NewToken(tokenString, tk)
}
