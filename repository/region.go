package repository

import (
	"encoding/json"
	"strconv"

	"github.com/danielbintar/qwe-server/config"
	region_config "github.com/danielbintar/qwe-server/config/region"
	"github.com/danielbintar/qwe-server/model"
)

func regionUsersKey(id uint) string {
	return "regions:" + strconv.FormatUint(uint64(id), 10) + ":users"
}

func GetRegionCharactersPosition(id uint) []*model.CharacterPosition {
	var positions []*model.CharacterPosition
	r, err := config.RedisInstance().HGetAll(regionUsersKey(id)).Result()
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

func GetRegionCharacterPosition(id uint, characterID uint) *model.CharacterPosition {
	var position *model.CharacterPosition
	r, err := config.RedisInstance().HGet(regionUsersKey(id), strconv.FormatUint(uint64(characterID), 10)).Result()
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

func SetRegionCharacterPosition(regionID uint, pos model.CharacterPosition) {
	coordinate := map[string]uint{
		"x": pos.X,
		"y": pos.Y,
	}

	coordinateJson, _ := json.Marshal(coordinate)

	err := config.RedisInstance().HSet(regionUsersKey(regionID), strconv.FormatUint(uint64(pos.ID), 10), coordinateJson).Err()
	if err != nil { panic(err) }
}

func FindRegion(id uint) *model.Region {
	for _, region := range region_config.Instance().Regions {
		if region.ID == id {
			return region
		}
	}

	return nil
}
