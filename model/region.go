package model

import "net/http"

type Region struct {
	ID                 uint                 `yaml:"id"        json:"id"`
	Name               string               `yaml:"name"      json:"name"`
	CharactersPosition []*CharacterPosition `                 json:"characters"`
	RegionTowns        []*RegionTown        `yaml:"towns"     json:"towns"`
}

type RegionTown struct {
	ID       uint      `yaml:"id"       json:"id"`
	Position *position `yaml:"position" json:"position"`
}

func (self *Region) FindTownPosition(ID uint) *position {
	for _, regionTown := range self.RegionTowns {
		if regionTown.ID == ID {
			return regionTown.Position
		}
	}

	panic("not found")
}

func (self *Region) Render(w http.ResponseWriter, r *http.Request) error {
	if self.CharactersPosition == nil {
		self.CharactersPosition = []*CharacterPosition{}
	}

	return nil
}
