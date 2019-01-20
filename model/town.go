package model

import "net/http"

type Town struct {
	Id   int    `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
}

func (f *Town) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

