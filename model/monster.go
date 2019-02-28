package model

type Monster struct {
	ID   uint   `yaml:"id"   json:"id"`
	Name string `yaml:"name" json:"name"`
}

type MonsterSpawn struct {
	ID        uint     `json:"id"`
	MonsterID uint     `json:"monster_id"`
	RegionID  uint     `json:"region_id"`
	Position  Position `json:"position"`
}
