package repository

import (
	"encoding/json"
	"strconv"

	"github.com/danielbintar/qwe-server/config"
)

func currentCharacterKey() string {
	return "current_character"
}

func characterTownIDKey() string {
	return "character_town_id"
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

func SetCurrentCharacter(userID uint, characterID uint) {
	err := config.RedisInstance().HSet(currentCharacterKey(), strconv.FormatUint(uint64(userID), 10), characterID).Err()
	if err != nil { panic(err) }
}

func GetCharacterTownID(id uint) *uint {
	var townID uint

	r, err := config.RedisInstance().HGet(characterTownIDKey(), strconv.FormatUint(uint64(id), 10)).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil
		} else {
			panic(err)
		}
	} else {
		json.Unmarshal([]byte(r), &townID)
	}

	return &townID
}

func SetCharacterTownID(characterID uint, townID uint) {
	err := config.RedisInstance().HSet(characterTownIDKey(), strconv.FormatUint(uint64(characterID), 10), townID).Err()
	if err != nil { panic(err) }
}
