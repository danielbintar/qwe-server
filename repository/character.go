package repository

import (
	"encoding/json"
	"strconv"

	"github.com/danielbintar/qwe-server/config"
)

func currentCharacterKey() string {
	return "current_character"
}

func characterTownKey() string {
	return "character_town"
}

func characterRegionKey() string {
	return "character_region"
}

func GetPlayingCharacter(id uint) *uint {
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

func SetPlayingCharacter(userID uint, characterID uint) {
	err := config.RedisInstance().HSet(currentCharacterKey(), strconv.FormatUint(uint64(userID), 10), characterID).Err()
	if err != nil { panic(err) }
}

func GetCharacterInTown(id uint) *uint {
	var townID uint

	r, err := config.RedisInstance().HGet(characterTownKey(), strconv.FormatUint(uint64(id), 10)).Result()
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

func SetCharacterInTown(characterID uint, townID uint) {
	err := config.RedisInstance().HSet(characterTownKey(), strconv.FormatUint(uint64(characterID), 10), townID).Err()
	if err != nil { panic(err) }
}

func SetCharacterInRegion(characterID uint, regionID uint) {
	err := config.RedisInstance().HSet(characterRegionKey(), strconv.FormatUint(uint64(characterID), 10), regionID).Err()
	if err != nil { panic(err) }
}

func UnsetCharacterInTown(characterID uint) {
	err := config.RedisInstance().HDel(characterTownKey(), strconv.FormatUint(uint64(characterID), 10)).Err()
	if err != nil { panic(err) }
}

func SetCurrentCharacter(userID uint, characterID uint) {
	err := config.RedisInstance().HSet(currentCharacterKey(), strconv.FormatUint(uint64(userID), 10), characterID).Err()
	if err != nil { panic(err) }

	if !characterHasPosition(characterID) {
		setDefaultPosition(characterID)
	}
}

func characterHasPosition(characterID uint) bool {
	return getTownCharacterPosition(uint(1), characterID) != nil
}
