package model

import "net/http"

type Town struct {
	Id    int            `yaml:"id"   json:"id"`
	Name  string         `yaml:"name" json:"name"`
	Users []UserPosition `            json:"users"`
}

func (f *Town) Render(w http.ResponseWriter, r *http.Request) error {
	if f.Users == nil {
		f.Users = []UserPosition{}
	}

	return nil
}
