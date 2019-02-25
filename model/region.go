package model

import "net/http"

type regionTown struct {
	ID            uint          `yaml:"id"             json:"id"`
	SpawnPosition position      `yaml:"spawn_position" json:"spawn_position"`
	Portals       rangePosition `yaml:"portal"         json:"portal"`
}

type Region struct {
	ID                 uint                 `yaml:"id"    json:"id"`
	Name               string               `yaml:"name"  json:"name"`
	Towns              []*regionTown        `yaml:"towns" json:"towns"`
	CharactersPosition []*CharacterPosition `             json:"characters"`
}

func (f *Region) Render(w http.ResponseWriter, r *http.Request) error {
	if f.CharactersPosition == nil {
		f.CharactersPosition = []*CharacterPosition{}
	}

	return nil
}
