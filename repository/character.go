package repository

import (
	"encoding/json"
	"strconv"

	"github.com/danielbintar/qwe-server/config"
	"github.com/danielbintar/qwe-server/model"
)

func currentCharacterKey() string {
	return "current_character"
}

func GetCurrentCharacter(id uint) *uint {
	var characterID uint
	r, err := config.RedisInstance().HGet(currentCharacterKey(), strconv.FormatUint(uint64(id), 10)).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil
		} else {
			panic(err)
		}
	} else {
		json.Unmarshal([]byte(r), &characterID)
	}

	return &characterID
}

func SetCurrentCharacter(userID uint, characterID uint) *model.CharacterPosition {
	err := config.RedisInstance().HSet(currentCharacterKey(), strconv.FormatUint(uint64(userID), 10), characterID).Err()
	if err != nil { panic(err) }

	position := getTownCharacterPosition(uint(1), characterID)
	if position == nil {
		position = setDefaultPosition(characterID)
	}

	return position
}
