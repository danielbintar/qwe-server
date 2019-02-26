package model

import "net/http"

type regionTown struct {
	ID            uint          `yaml:"id"             json:"id"`
	SpawnPosition position      `yaml:"spawn_position" json:"spawn_position"`
	Portal        rangePosition `yaml:"portal"         json:"portal"`
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

func (self Region) FindTown(id uint) *regionTown {
	for _, town := range self.Towns {
		if town.ID == id {
			return town
		}
	}
	return nil
}
