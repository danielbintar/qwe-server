package repository

import (
	"encoding/json"
	"strconv"

	"github.com/danielbintar/qwe-server/config"
	town_config "github.com/danielbintar/qwe-server/config/town"
	"github.com/danielbintar/qwe-server/model"
)

func townUsersKey(id uint) string {
	return "towns:" + strconv.FormatUint(uint64(id), 10) + ":users"
}

func GetTownCharactersPosition(id uint) []*model.CharacterPosition {
	var positions []*model.CharacterPosition
	r, err := config.RedisInstance().HGetAll(townUsersKey(id)).Result()
	if err != nil {
		if err.Error() != "redis: nil" {
			panic(err)
		}
	} else {
		for k, v := range r {
			var position *model.CharacterPosition
			json.Unmarshal([]byte(v), &position)
			u64, _ := strconv.ParseUint(k, 10, 32)
			position.ID = uint(u64)
			positions = append(positions, position)
		}
	}
	return positions
}

func SetTownCharacterPosition(townID uint, position *model.CharacterPosition) {
	coordinate := map[string]uint{
		"x": position.X,
		"y": position.Y,
	}

	coordinateJson, _ := json.Marshal(coordinate)

	err := config.RedisInstance().HSet(townUsersKey(townID), strconv.FormatUint(uint64(position.ID), 10), coordinateJson).Err()
	if err != nil { panic(err) }
}


func FindTown(id uint) *model.Town {
	for _, town := range town_config.Instance().Towns {
		if town.ID == id {
			return town
		}
	}

	return &model.Town{}
}
