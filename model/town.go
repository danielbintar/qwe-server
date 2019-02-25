package model

import "net/http"

type position struct {
	X uint `yaml:"x" json:"x"`
	Y uint `yaml:"y" json:"y"`
}

type rangePosition struct {
	MinX uint `yaml:"min_x" json:"min_x"`
	MaxX uint `yaml:"max_x" json:"max_x"`
	MinY uint `yaml:"min_y" json:"min_y"`
	MaxY uint `yaml:"max_y" json:"max_y"`
}

func (self rangePosition) In(p CharacterPosition) bool {
	return p.X >= self.MinX && p.X <= self.MaxX &&
		p.Y >= self.MinY && p.Y <= self.MaxY
}

type Town struct {
	ID                 uint                 `yaml:"id"       json:"id"`
	Name               string               `yaml:"name"     json:"name"`
	Position           position             `yaml:"position" json:"position"`
	Portals            []*rangePosition     `yaml:"portals"  json:"portals"`
	CharactersPosition []*CharacterPosition `                json:"characters"`
}

func (f *Town) Render(w http.ResponseWriter, r *http.Request) error {
	if f.CharactersPosition == nil {
		f.CharactersPosition = []*CharacterPosition{}
	}

	return nil
}
