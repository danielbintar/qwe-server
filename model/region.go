package model

import "net/http"

type regionTown struct {
	ID            uint          `yaml:"id"             json:"id"`
	SpawnPosition Position      `yaml:"spawn_position" json:"spawn_position"`
	Portal        RangePosition `yaml:"portal"         json:"portal"`
}

type regionMonster struct {
	ID       uint          `yaml:"id"       json:"id"`
	Total    uint          `yaml:"total"    json:"total"`
	Position RangePosition `yaml:"position" json:"position"`
}

type Region struct {
	ID                 uint                 `yaml:"id"    json:"id"`
	Name               string               `yaml:"name"  json:"name"`
	Towns              []*regionTown        `yaml:"towns" json:"towns"`
	Monsters           []*regionMonster     `yaml:"monsters"`
	CharactersPosition []*CharacterPosition `             json:"characters"`
}

func (self Region) Serialize() *RegionSerializer {
	return &RegionSerializer {
		ID: self.ID,
		Name: self.Name,
		Towns: self.Towns,
		CharactersPosition: self.CharactersPosition,
	}
}

type RegionSerializer struct {
	ID                 uint                 `json:"id"`
	Name               string               `json:"name"`
	Towns              []*regionTown        `json:"towns"`
	Monsters           []*MonsterSpawn      `json:"monsters"`
	CharactersPosition []*CharacterPosition `json:"characters"`
}

func (f *RegionSerializer) Render(w http.ResponseWriter, r *http.Request) error {
	if f.CharactersPosition == nil {
		f.CharactersPosition = []*CharacterPosition{}
	}

	if f.Monsters == nil {
		f.Monsters = []*MonsterSpawn{}
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
