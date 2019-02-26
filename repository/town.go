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

			if IsLoginCharacter(uint(u64)) {
				position.ID = uint(u64)
				positions = append(positions, position)
			}
		}
	}
	return positions
}

func GetTownCharacterPosition(id uint, characterID uint) *model.CharacterPosition {
	var position *model.CharacterPosition
	r, err := config.RedisInstance().HGet(townUsersKey(id), strconv.FormatUint(uint64(characterID), 10)).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil
		} else {
			panic(err)
		}
	} else {
		json.Unmarshal([]byte(r), &position)
		position.ID = characterID
	}
	return position
}

func SetTownCharacterPosition(townID uint, pos model.CharacterPosition) {
	coordinate := map[string]uint{
		"x": pos.X,
		"y": pos.Y,
	}

	coordinateJson, _ := json.Marshal(coordinate)

	err := config.RedisInstance().HSet(townUsersKey(townID), strconv.FormatUint(uint64(pos.ID), 10), coordinateJson).Err()
	if err != nil { panic(err) }
}

func UnsetTownCharacterPosition(townID uint, characterID uint) {
	err := config.RedisInstance().HDel(townUsersKey(townID), strconv.FormatUint(uint64(characterID), 10)).Err()
	if err != nil { panic(err) }
}

func FindTown(id uint) *model.Town {
	for _, town := range town_config.Instance().Towns {
		if town.ID == id {
			return town
		}
	}

	return nil
}
