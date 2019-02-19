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

func UnsetTownCharacterPosition(characterID uint, townID uint) {
	err := config.RedisInstance().HDel(townUsersKey(townID), strconv.FormatUint(uint64(characterID), 10)).Err()
	if err != nil { panic(err) }
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

func getTownCharacterPosition(id uint, characterID uint) *model.CharacterPosition {
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

func SetTownCharacterPosition(id uint, position *model.CharacterPosition) {
	coordinate := map[string]uint{
		"x": position.X,
		"y": position.Y,
	}

	coordinateJson, _ := json.Marshal(coordinate)
	err := config.RedisInstance().HSet(townUsersKey(id), strconv.FormatUint(uint64(position.ID), 10), coordinateJson).Err()
	if err != nil { panic(err) }
}

func setDefaultPosition(characterID uint) {
	townID := uint(1)
	town := FindTown(townID)

	coordinate := map[string]uint{
		"x": town.Position.X,
		"y": town.Position.Y,
	}

	coordinateJson, _ := json.Marshal(coordinate)

	err := config.RedisInstance().HSet(townUsersKey(townID), strconv.FormatUint(uint64(characterID), 10), coordinateJson).Err()
	if err != nil { panic(err) }
}

func MovingCharacter(townID uint, movement *model.CharacterMovement) *model.CharacterPosition {
	position := getTownCharacterPosition(townID, movement.ID)

	if int(position.X) + movement.X > 0 {
		position.X = uint(int(position.X) + movement.X)
	} else {
		position.X = 0
	}

	if int(position.Y) + movement.Y > 0 {
		position.Y = uint(int(position.Y) + movement.Y)
	} else {
		position.Y = 0
	}

	coordinate := map[string]uint{
		"x": position.X,
		"y": position.Y,
	}

	coordinateJson, _ := json.Marshal(coordinate)

	err := config.RedisInstance().HSet(townUsersKey(townID), strconv.FormatUint(uint64(position.ID), 10), coordinateJson).Err()
	if err != nil { panic(err) }
	return position
}


func FindTown(id uint) *model.Town {
	for _, town := range town_config.Instance().Towns {
		if town.ID == id {
			return town
		}
	}

	return &model.Town{}
}
