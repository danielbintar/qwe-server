package web

import (
	"os"
	"strings"
	"context"
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

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		//The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		//Grab the token part, what we are truly interested in
		tokenPart := splitted[1]
		tk := &model.Jwt{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("APPLICATION_SECRET_KEY")), nil
		})

		// Malformed token
		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		// Token is invalid, maybe not signed on this server
		if !token.Valid {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		ctx := context.WithValue(r.Context(), "jwt", tk)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
