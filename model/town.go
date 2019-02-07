package model

import "net/http"

type position struct {
	X uint `yml:"x" json:"x"`
	Y uint `yml:"y" json:"y"`
}

type Town struct {
	ID       uint           `yaml:"id"        json:"id"`
	Name     string         `yaml:"name"      json:"name"`
	Position position       `yaml:"position"  json:"position"`
	Users    []UserPosition `                 json:"users"`
}

func (f *Town) Render(w http.ResponseWriter, r *http.Request) error {
	if f.Users == nil {
		f.Users = []UserPosition{}
	}

	return nil
}
