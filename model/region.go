package model

import "net/http"

type Region struct {
	ID                 uint                 `yaml:"id"        json:"id"`
	Name               string               `yaml:"name"      json:"name"`
	CharactersPosition []*CharacterPosition `                 json:"characters"`
	RegionTowns        *RegionTowns         `yaml:"towns"     json:"towns"`
}

type RegionTowns struct {
	ID       uint     `yaml:"id"       json:"id"`
	Position position `yaml:"position" json:"position"`
}

func (f *Region) Render(w http.ResponseWriter, r *http.Request) error {
	if f.CharactersPosition == nil {
		f.CharactersPosition = []*CharacterPosition{}
	}

	return nil
}
